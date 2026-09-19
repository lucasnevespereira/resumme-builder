package render

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io/fs"
	"path"
	"strings"

	"github.com/lucasnevespereira/resumme-builder/internal/models"
)

const DefaultTheme = "classic"

// Printer turns a rendered HTML page into PDF bytes.
type Printer interface {
	Print(ctx context.Context, html []byte) ([]byte, error)
}

type Renderer struct {
	themes  fs.FS
	labels  map[string]map[string]string
	printer Printer
}

// New reads themes from templates/ and labels from locales/ inside ui.
func New(ui fs.FS, printer Printer) (*Renderer, error) {
	labels, err := loadLabels(ui)
	if err != nil {
		return nil, err
	}
	themes, err := fs.Sub(ui, "templates")
	if err != nil {
		return nil, err
	}
	return &Renderer{themes: themes, labels: labels, printer: printer}, nil
}

// page is what themes see. It has the resume fields at the root, plus Labels.
type page struct {
	models.Resume
	Labels map[string]string
}

func (r *Renderer) HTML(resume models.Resume) ([]byte, error) {
	theme := resume.Meta.Template
	if theme == "" {
		theme = DefaultTheme
	}

	t, err := r.theme(theme)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, page{Resume: resume, Labels: r.labelsFor(resume.Meta.Lang)}); err != nil {
		return nil, fmt.Errorf("render %s: %w", theme, err)
	}
	return buf.Bytes(), nil
}

func (r *Renderer) PDF(ctx context.Context, resume models.Resume) ([]byte, error) {
	html, err := r.HTML(resume)
	if err != nil {
		return nil, err
	}
	return r.printer.Print(ctx, html)
}

func (r *Renderer) theme(name string) (*template.Template, error) {
	// The name comes from request data, so it must stay a single directory.
	if strings.ContainsAny(name, `/\*?[`) || name == "." || name == ".." {
		return nil, fmt.Errorf("unknown template %q", name)
	}
	files, err := fs.Glob(r.themes, name+"/*")
	if err != nil || len(files) == 0 {
		return nil, fmt.Errorf("unknown template %q", name)
	}

	// Themes name their root file _<theme> so it sorts first and is the one executed.
	t := template.New(path.Base(files[0])).Funcs(funcs).Option("missingkey=error")
	t, err = t.ParseFS(r.themes, files...)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", name, err)
	}
	return t, nil
}
