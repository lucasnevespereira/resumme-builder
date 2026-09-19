package chrome

import (
	"context"
	"math"
	"os"
	"resumme-builder/internal/utils/logger"
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
	if err := chromedp.Run(chromeCtx, printTasks("file://"+file.Name(), &pdfData)); err != nil {
		return nil, errors.Wrap(err, "Print - chromedp.Run")
	}

	logger.Log.Infof("PDF generated in %f seconds", time.Since(startedAt).Seconds())

	return pdfData, nil
}

func printTasks(url string, pdf *[]byte) chromedp.Tasks {
	idle := make(chan struct{}, 1)
	return chromedp.Tasks{
		emulation.SetUserAgentOverride(userAgentOverride),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// A page target's main frame shares its ID. Iframes fire their own
			// networkIdle, so only the main frame counts, and repeats are dropped.
			mainFrame := cdp.FrameID(chromedp.FromContext(ctx).Target.TargetID)
			chromedp.ListenTarget(ctx, func(ev interface{}) {
				if e, ok := ev.(*page.EventLifecycleEvent); ok && e.Name == "networkIdle" && e.FrameID == mainFrame {
					select {
					case idle <- struct{}{}:
					default:
					}
				}
			})
			return nil
		}),
		chromedp.Navigate(url),
		chromedp.WaitVisible(htmlSelector, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			select {
			case <-idle:
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(networkIdleTimeout):
				// Slow fonts or images shouldn't fail the print.
				logger.Log.Warnf("no network idle after %s, printing anyway", networkIdleTimeout)
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
				return errors.Wrap(err, "printTasks - page.PrintToPDF")
			}
			*pdf = data
			return nil
		}),
	}
}
