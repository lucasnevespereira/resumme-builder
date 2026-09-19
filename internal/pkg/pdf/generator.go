package pdf

import (
	"context"
	"fmt"
	"math"
	"os"
	"resumme-builder/internal/utils/logger"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
)

const (
	userAgentOverride   = "WebScraper 1.0"
	htmlSelector        = "body"
	networkReadyTimeOut = 15 * time.Second

	paperWidthInches  = 8.3
	paperHeightInches = 11.7
	cssPixelsPerInch  = 96
)

var (
	paperWidthPx  = int64(math.Round(paperWidthInches * cssPixelsPerInch))
	paperHeightPx = int64(math.Round(paperHeightInches * cssPixelsPerInch))
)

// Generator provides functionality to generate PDF from HTML.
type Generator struct{}

// NewPDFGenerator creates a new instance of PDFGenerator.
func NewPDFGenerator() *Generator {
	return &Generator{}
}

// Print loads html in headless Chrome and prints it to PDF.
func (g *Generator) Print(ctx context.Context, html []byte) ([]byte, error) {
	startedAt := time.Now()

	// Each call gets its own file, so concurrent requests never share one.
	file, err := os.CreateTemp("", "resume-*.html")
	if err != nil {
		return nil, errors.Wrap(err, "Print - os.CreateTemp")
	}
	defer os.Remove(file.Name())
	if _, err := file.Write(html); err != nil {
		file.Close()
		return nil, errors.Wrap(err, "Print - write html")
	}
	if err := file.Close(); err != nil {
		return nil, errors.Wrap(err, "Print - close html")
	}

	chromeCtx, cancelCtx := chromedp.NewContext(ctx)
	defer cancelCtx()

	var pdfData []byte
	if err := chromedp.Run(chromeCtx, g.saveURLAsPDF("file://"+file.Name(), &pdfData)); err != nil {
		return nil, errors.Wrap(err, "Print - chromedp.Run")
	}

	logger.Log.Infof("PDF generated in %f seconds", time.Since(startedAt).Seconds())

	return pdfData, nil
}

func (g *Generator) saveURLAsPDF(url string, pdf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		emulation.SetUserAgentOverride(userAgentOverride),
		chromedp.Navigate(url),
		chromedp.WaitVisible(htmlSelector, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			if err := waitForNetworkIdle(ctx, networkReadyTimeOut); err != nil {
				logger.Log.Warn(err)
			}
			return nil
		}),
		// The template works out where the page breaks fall itself. That
		// measurement must run under output conditions: some widths depend on
		// vw units, hence on the paper size rather than on the window.
		emulation.SetDeviceMetricsOverride(paperWidthPx, paperHeightPx, 1, false),
		emulation.SetEmulatedMedia().WithMedia("print"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.Evaluate("window.fitSidebar && window.fitSidebar()", nil).Do(ctx)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			data, _, err := page.
				PrintToPDF().
				WithMarginLeft(0).
				WithMarginTop(0).
				WithMarginRight(0).
				WithMarginBottom(0).
				WithPaperWidth(paperWidthInches).
				WithPaperHeight(paperHeightInches).
				WithPrintBackground(true).
				Do(ctx)
			if err != nil {
				return errors.Wrap(err, "saveURLAsPDF - page.PrintToPDF")
			}
			*pdf = data
			return nil
		}),
	}
}

func waitForNetworkIdle(ctx context.Context, timeout time.Duration) error {
	idleChan := make(chan struct{})

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if event, ok := ev.(*page.EventLifecycleEvent); ok {
			if event.Name == "networkIdle" {
				close(idleChan)
			}
		}
	})

	select {
	case <-idleChan:
		// Network is idle
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("timeout %.0f seconds waiting for network idle", timeout.Seconds())
	}
}
