package pdf

import (
	"context"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"os"
	"resumme-builder/internal/utils/logger"
	"time"
)

const (
	userAgentOverride = "WebScraper 1.0"
	htmlSelector      = "body"
)

func Write(destination string, data []byte) error {
	if err := os.WriteFile(destination, data, os.ModePerm); err != nil {
		return errors.Wrap(err, "pdf.Write")
	}

	return nil
}

func GenerateFromHtml(file string) ([]byte, error) {
	startedAt := time.Now()

	chromeCtx, cancelCtx := chromedp.NewContext(context.Background())
	defer cancelCtx()

	var pdfData []byte
	url := getFilePathAsUrl(file)

	if err := chromedp.Run(chromeCtx, saveUrlAsPdf(url, &pdfData)); err != nil {
		return nil, errors.Wrap(err, "chromedp.Run")
	}

	logger.Log.Infof("Pdf generated in %f seconds\n", time.Since(startedAt).Seconds())

	return pdfData, nil
}

func saveUrlAsPdf(url string, pdf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		emulation.SetUserAgentOverride(userAgentOverride),
		chromedp.Navigate(url),
		chromedp.WaitVisible(htmlSelector, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			data, _, err := page.PrintToPDF().WithPrintBackground(true).Do(ctx)
			if err != nil {
				return errors.Wrap(err, "page.PrintToPDF")
			}
			*pdf = data
			return nil
		}),
	}
}

func getFilePathAsUrl(filename string) string {
	path, err := os.Getwd()
	if err != nil {
		logger.Log.Error(err)
	}

	return "file://" + path + "/" + filename
}
