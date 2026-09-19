package chrome

import (
	"context"
	"fmt"
	"github.com/lucasnevespereira/resumme-builder/internal/utils/logger"
	"math"
	"os"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
)

const (
	userAgentOverride  = "WebScraper 1.0"
	htmlSelector       = "body"
	networkIdleTimeout = 15 * time.Second
	printTimeout       = time.Minute

	paperWidthInches  = 8.3
	paperHeightInches = 11.7
	cssPixelsPerInch  = 96
)

var (
	paperWidthPx  = int64(math.Round(paperWidthInches * cssPixelsPerInch))
	paperHeightPx = int64(math.Round(paperHeightInches * cssPixelsPerInch))
)

// Printer prints HTML to PDF in headless Chrome.
type Printer struct{}

func NewPrinter() *Printer {
	return &Printer{}
}

// Print loads html in a fresh Chrome and prints it to PDF.
func (p *Printer) Print(ctx context.Context, html []byte) ([]byte, error) {
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

	ctx, cancel := context.WithTimeout(ctx, printTimeout)
	defer cancel()
	chromeCtx, cancelCtx := chromedp.NewContext(ctx)
	defer cancelCtx()

	var pdfData []byte
	if err := chromedp.Run(chromeCtx, p.saveURLAsPDF("file://"+file.Name(), &pdfData)); err != nil {
		return nil, errors.Wrap(err, "Print - chromedp.Run")
	}

	logger.Log.Infof("PDF generated in %f seconds", time.Since(startedAt).Seconds())

	return pdfData, nil
}

func (p *Printer) saveURLAsPDF(url string, pdf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		emulation.SetUserAgentOverride(userAgentOverride),
		chromedp.Navigate(url),
		chromedp.WaitVisible(htmlSelector, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			if err := waitForNetworkIdle(ctx, networkIdleTimeout); err != nil {
				if ctx.Err() != nil {
					return err
				}
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

// waitForNetworkIdle waits until the page has had no network activity for
// 500ms, so fonts and icons loaded late still make it into the PDF.
func waitForNetworkIdle(ctx context.Context, timeout time.Duration) error {
	mainFrame := cdp.FrameID(chromedp.FromContext(ctx).Target.TargetID)
	idle := make(chan struct{}, 1)

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		// Iframes fire their own networkIdle, and Chrome can fire it more than
		// once. Only the main frame counts, and repeats are dropped.
		event, ok := ev.(*page.EventLifecycleEvent)
		if !ok || event.Name != "networkIdle" || event.FrameID != mainFrame {
			return
		}
		select {
		case idle <- struct{}{}:
		default:
		}
	})

	select {
	case <-idle:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(timeout):
		return fmt.Errorf("timeout %.0f seconds waiting for network idle", timeout.Seconds())
	}
}
