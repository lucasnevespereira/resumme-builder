package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"resumme-builder/internal/render"
)

type fakePrinter struct{}

func (fakePrinter) Print(context.Context, []byte) ([]byte, error) {
	return []byte("%PDF"), nil
}

func postPdf(t *testing.T, body string) *httptest.ResponseRecorder {
	t.Helper()
	renderer, err := render.New(os.DirFS("../../../ui"), fakePrinter{})
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

func TestGetPdfRejectsInvalidJSON(t *testing.T) {
	if rec := postPdf(t, "{"); rec.Code != http.StatusBadRequest {
		t.Errorf("got %d, want 400", rec.Code)
	}
}

// Used to call logger.Fatal, which exits the process: this test would not return.
func TestGetPdfReportsRenderErrors(t *testing.T) {
	rec := postPdf(t, `{"meta": {"template": "nope"}}`)
	if rec.Code != http.StatusInternalServerError || !strings.Contains(rec.Body.String(), "unknown template") {
		t.Errorf("got %d %s", rec.Code, rec.Body.String())
	}
}
