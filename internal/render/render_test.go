package render

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"os"
	"strings"
	"testing"

	"resumme-builder/internal/models"
)

func newTestRenderer(t *testing.T) *Renderer {
	t.Helper()
	r, err := New(os.DirFS("../../ui"), nil)
	if err != nil {
		t.Fatal(err)
	}
	return r
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
	r := newTestRenderer(t)
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
	r := newTestRenderer(t)
	resume := exampleResume(t)
	resume.Meta.Template = "stackoverflow"

	for lang, want := range map[string]string{"fr": "Formation", "fr_FR": "Formation", "de": "Education", "": "Education"} {
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

func TestUnknownTemplateIsAnError(t *testing.T) {
	r := newTestRenderer(t)
	resume := exampleResume(t)

	for _, name := range []string{"nope", "..", "../locales", "classic/..", "*"} {
		resume.Meta.Template = name
		if _, err := r.HTML(resume); err == nil || !strings.Contains(err.Error(), "unknown template") {
			t.Errorf("template %q: want unknown template error, got %v", name, err)
		}
	}
}
