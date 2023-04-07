package pdf

import (
	"context"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"resumme-builder/internal/utils"
	"time"
)

func Generate(url string) ([]byte, error) {
	var pdfData []byte

	chromeCtx, cancelCtx := chromedp.NewContext(context.Background())
	defer cancelCtx()

	startedAt := time.Now()
	if err := chromedp.Run(chromeCtx, saveUrlAsPdf(url, &pdfData)); err != nil {
		return nil, errors.Wrap(err, "GeneratePDF - chromedp.Run")
	}

	utils.Logger.Infof("Pdf generated in %f seconds\n", time.Since(startedAt).Seconds())

	return pdfData, nil
}

func saveUrlAsPdf(url string, pdf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		emulation.SetUserAgentOverride("WebScraper 1.0"),
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			data, _, err := page.PrintToPDF().Do(ctx)
			if err != nil {
				return errors.Wrap(err, "page.PrintToPDF")
			}
			*pdf = data
			return nil
		}),
	}
}
