package template

import (
	"github.com/pkg/errors"
	"html/template"
	"path"
	"path/filepath"
)

type Manager struct {
	TemplateDir string
}

func NewTemplateManager(templateDir string) *Manager {
	return &Manager{
		TemplateDir: templateDir,
	}
}

func (tm *Manager) GetTemplate(name string) (*template.Template, error) {
	templateFiles, err := filepath.Glob(filepath.Join(tm.TemplateDir, name, "*"))
	if err != nil {
		return nil, errors.Wrap(err, "GetTemplate filepath.Glob")
	}

	templateFuncs := template.FuncMap{
		"isLast":          isLast,
		"displayLocation": displayLocation,
		"trimURLPrefix":   trimURLPrefix,
		"getFirstName":    getFirstName,
		"getLastName":     getLastName,
		"evaluate":        evaluate,
		"lowerEq":         lowerEq,
	}

	t := template.New(path.Base(templateFiles[0])).Funcs(templateFuncs)
	t, err = t.ParseFiles(templateFiles...)
	if err != nil {
		return nil, errors.Wrap(err, "GetTemplate template.New")
	}

	return t, nil
}
