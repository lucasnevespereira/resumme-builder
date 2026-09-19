package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/lucasnevespereira/resb/internal/render"
	"github.com/lucasnevespereira/resb/ui"
)

type fakePrinter struct{}

func (fakePrinter) Print(context.Context, []byte) ([]byte, error) {
	return []byte("%PDF"), nil
}

func postPdf(t *testing.T, body string) *httptest.ResponseRecorder {
	t.Helper()
	renderer, err := render.New(ui.FS, fakePrinter{})
	if err != nil {
		t.Fatal(err)
	}
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/pdf", GetPdf(renderer))

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/pdf", strings.NewReader(body)))
	return rec
}

func TestGetPdfReturnsThePdf(t *testing.T) {
	body, err := os.ReadFile("../../../examples/example.resume.json")
	if err != nil {
		t.Fatal(err)
	}
	rec := postPdf(t, string(body))

	if rec.Code != http.StatusOK || rec.Header().Get("Content-Type") != "application/pdf" || rec.Body.String() != "%PDF" {
		t.Errorf("got %d %q %q", rec.Code, rec.Header().Get("Content-Type"), rec.Body.String())
	}
}

// The handler used to call logger.Fatal here, which exits the process.
func TestGetPdfReportsRenderErrors(t *testing.T) {
	rec := postPdf(t, `{"meta": {"template": "nope"}}`)
	if rec.Code != http.StatusInternalServerError || !strings.Contains(rec.Body.String(), "unknown template") {
		t.Errorf("got %d %s", rec.Code, rec.Body.String())
	}
}
