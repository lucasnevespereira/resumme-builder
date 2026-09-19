package chrome

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
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

// Icon kits and web fonts are often fetched by a script after the load
// event. The print has to wait for them, not just for the HTML.
func TestPrintWaitsForLateSlowAssets(t *testing.T) {
	if testing.Short() {
		t.Skip("needs Chrome")
	}
	const delay = 2 * time.Second
	var served atomic.Int64
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.Header().Set("Content-Type", "image/svg+xml")
		fmt.Fprint(w, `<svg xmlns="http://www.w3.org/2000/svg" width="10" height="10"/>`)
		served.Store(time.Now().UnixNano())
	}))
	defer srv.Close()

	html := fmt.Sprintf(`<html><body>x<script>
window.addEventListener("load", () => setTimeout(() => {
  const img = new Image(); img.src = %q; document.body.appendChild(img);
}, 100));
</script></body></html>`, srv.URL+"/icon.svg")

	if _, err := NewPrinter().Print(context.Background(), []byte(html)); err != nil {
		t.Fatal(err)
	}
	printed := time.Now().UnixNano()
	if served.Load() == 0 || served.Load() > printed {
		t.Error("printed before the late asset arrived")
	}
}
