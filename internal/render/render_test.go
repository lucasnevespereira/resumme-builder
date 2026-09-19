package render

import (
	"bytes"
	"context"
	"encoding/json"
	"io/fs"
	"os"
	"strings"
	"testing"

	"resumme-builder/internal/models"
)

type fakePrinter struct{ got []byte }

func (p *fakePrinter) Print(_ context.Context, html []byte) ([]byte, error) {
	p.got = html
	return []byte("%PDF"), nil
}

func newTestRenderer(t *testing.T) (*Renderer, *fakePrinter) {
	t.Helper()
	printer := &fakePrinter{}
	r, err := New(os.DirFS("../../ui"), printer)
	if err != nil {
		t.Fatal(err)
	}
	return r, printer
}

func exampleResume(t *testing.T) models.Resume {
	t.Helper()
	data, err := os.ReadFile("../../examples/example.resume.json")
	if err != nil {
		t.Fatal(err)
	}
	var resume models.Resume
	if err := json.Unmarshal(data, &resume); err != nil {
		t.Fatal(err)
	}
	return resume
}

func TestEveryThemeRendersTheExample(t *testing.T) {
	r, _ := newTestRenderer(t)
	themes, err := fs.ReadDir(os.DirFS("../../ui/templates"), ".")
	if err != nil {
		t.Fatal(err)
	}

	for _, theme := range themes {
		for _, lang := range []string{"en", "fr"} {
			t.Run(theme.Name()+"/"+lang, func(t *testing.T) {
				resume := exampleResume(t)
				resume.Meta.Template = theme.Name()
				resume.Meta.Lang = lang

				html, err := r.HTML(resume)
				if err != nil {
					t.Fatal(err)
				}
				if !bytes.Contains(html, []byte("John Doe")) {
					t.Error("rendered page is missing the resume name")
				}
			})
		}
	}
}

func TestLabelsFollowTheResumeLanguage(t *testing.T) {
	r, _ := newTestRenderer(t)
	resume := exampleResume(t)
	resume.Meta.Template = "stackoverflow"

	for lang, want := range map[string]string{"fr": "Formation", "fr_FR": "Formation", "fr-FR": "Formation", "de": "Education", "": "Education"} {
		resume.Meta.Lang = lang
		html, err := r.HTML(resume)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Contains(html, []byte(want)) {
			t.Errorf("lang %q: want label %q", lang, want)
		}
	}
}

func TestEveryLocaleHasTheEnglishKeys(t *testing.T) {
	r, _ := newTestRenderer(t)
	for lang, labels := range r.labels {
		for key := range r.labels[defaultLang] {
			if _, ok := labels[key]; !ok {
				t.Errorf("locales/%s.json is missing %q", lang, key)
			}
		}
	}
}

func TestEmptyTemplateUsesTheDefaultTheme(t *testing.T) {
	r, _ := newTestRenderer(t)
	resume := exampleResume(t)

	resume.Meta.Template = ""
	got, err := r.HTML(resume)
	if err != nil {
		t.Fatal(err)
	}
	resume.Meta.Template = DefaultTheme
	want, _ := r.HTML(resume)
	if !bytes.Equal(got, want) {
		t.Error("empty template should render the default theme")
	}
}

func TestUnknownTemplateIsAnError(t *testing.T) {
	r, _ := newTestRenderer(t)
	resume := exampleResume(t)

	for _, name := range []string{"nope", "..", "../locales", "classic/..", "*"} {
		resume.Meta.Template = name
		if _, err := r.HTML(resume); err == nil || !strings.Contains(err.Error(), "unknown template") {
			t.Errorf("template %q: want unknown template error, got %v", name, err)
		}
	}
}

func TestPDFPrintsTheRenderedHTML(t *testing.T) {
	r, printer := newTestRenderer(t)
	resume := exampleResume(t)

	pdf, err := r.PDF(context.Background(), resume)
	if err != nil {
		t.Fatal(err)
	}
	html, _ := r.HTML(resume)
	if string(pdf) != "%PDF" || !bytes.Equal(printer.got, html) {
		t.Error("PDF should print exactly the rendered HTML")
	}
}

func TestDatesFollowTheSameLanguageAsLabels(t *testing.T) {
	for _, lang := range []string{"fr", "fr_FR", "fr-FR", "FR"} {
		if got := formatDate("January 2006", "2020-01-15", lang); got != "janvier 2020" {
			t.Errorf("lang %q: got %q, want janvier 2020", lang, got)
		}
	}
	if got := formatDate("January 2006", "2020-01-15", "de"); got != "January 2020" {
		t.Errorf("unknown lang: got %q, want English", got)
	}
}
