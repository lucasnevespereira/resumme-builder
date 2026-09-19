package chrome

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"
)

// Needs a local Chrome; skipped with -short or when Chrome is missing.
func TestPrintPageWithIframes(t *testing.T) {
	if testing.Short() {
		t.Skip("needs Chrome")
	}
	html := []byte(`<html><body><h1>Resume</h1>
<iframe srcdoc="<p>one</p>"></iframe><iframe srcdoc="<p>two</p>"></iframe>
</body></html>`)

	// Iframes fire their own networkIdle. This used to close a closed channel
	// and crash the process, and must not wait out networkIdleTimeout either.
	started := time.Now()
	pdf, err := NewPrinter().Print(context.Background(), html)
	if err != nil && strings.Contains(err.Error(), "executable file not found") {
		t.Skip("Chrome not installed")
	}
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(pdf, []byte("%PDF")) {
		t.Errorf("not a PDF (%d bytes)", len(pdf))
	}
	if elapsed := time.Since(started); elapsed >= networkIdleTimeout {
		t.Errorf("took %s: missed the main frame's networkIdle", elapsed)
	}
}

func TestPrintStopsWhenCancelled(t *testing.T) {
	if testing.Short() {
		t.Skip("needs Chrome")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	started := time.Now()
	if _, err := NewPrinter().Print(ctx, []byte("<html><body>x</body></html>")); err == nil {
		t.Fatal("want an error from a cancelled print")
	}
	if elapsed := time.Since(started); elapsed > 5*time.Second {
		t.Errorf("cancelled print took %s", elapsed)
	}
}
