package pdf

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"os"
	"resume-builder/utils"
	"time"
)

func Generate() error {
	url := utils.GetFilePathAsUrl(utils.OutputHtmlFile)
	var pdfData []byte

	chromeCtx, cancelCtx := chromedp.NewContext(context.Background())
	defer cancelCtx()

	startedAt := time.Now()
	if err := chromedp.Run(chromeCtx, saveUrlAsPdf(url, &pdfData)); err != nil {
		return errors.Wrap(err, "GeneratePDF - chromedp.Run")
	}

	if err := os.WriteFile(utils.OutputPdfFile, pdfData, 0644); err != nil {
		return errors.Wrap(err, "GeneratePDF - WriteFile")
	}

	fmt.Printf("Pdf generated in %f seconds\n", time.Since(startedAt).Seconds())

	return nil
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
