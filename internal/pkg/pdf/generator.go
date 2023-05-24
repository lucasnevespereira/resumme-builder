package pdf

import (
	"context"
	"fmt"
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

// Generator provides functionality to generate PDF from HTML.
type Generator struct{}

// NewPDFGenerator creates a new instance of PDFGenerator.
func NewPDFGenerator() *Generator {
	return &Generator{}
}

// GenerateFromHTML generates a PDF from HTML file.
func (g *Generator) GenerateFromHTML(file string) ([]byte, error) {
	startedAt := time.Now()

	chromeCtx, cancelCtx := chromedp.NewContext(context.Background())
	defer cancelCtx()

	var pdfData []byte
	url := getFilePathAsURL(file)

	if err := chromedp.Run(chromeCtx, g.saveURLAsPDF(url, &pdfData)); err != nil {
		return nil, errors.Wrap(err, "GenerateFromHTML - chromedp.Run")
	}

	logger.Log.Infof("PDF %s generated in %f seconds", file, time.Since(startedAt).Seconds())

	return pdfData, nil
}

func (g *Generator) saveURLAsPDF(url string, pdf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		emulation.SetUserAgentOverride(userAgentOverride),
		chromedp.Navigate(url),
		chromedp.WaitVisible(htmlSelector, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			data, _, err := page.
				PrintToPDF().
				WithMarginLeft(0).
				WithMarginTop(0).
				WithMarginRight(0).
				WithMarginBottom(0).
				WithPaperWidth(8.3).
				WithPaperHeight(11.7).
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

func getFilePathAsURL(filename string) string {
	path, err := os.Getwd()
	if err != nil {
		logger.Log.Error(err)
	}

	return fmt.Sprintf("file://%s/%s", path, filename)
}
