package resume

import (
	"github.com/pkg/errors"
	"html/template"
	"path"
	"path/filepath"
	"resumme-builder/internal/models"
	"strings"
)

func getTemplate(name string) (*template.Template, error) {
	templateFiles, err := filepath.Glob("ui/" + name + "/*")
	if err != nil {
		return nil, errors.Wrap(err, "getTemplate filepath.Glob")
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
		return nil, errors.Wrap(err, "getTemplate template.New")
	}

	return t, nil
}

func isLast(index, length int) bool {
	return index == length-1
}

func displayLocation(location models.Location) string {
	var parts []string

	if location.City != "" {
		parts = append(parts, location.City)
	}
	if location.Region != "" {
		parts = append(parts, location.Region)
	}
	if location.CountryCode != "" {
		parts = append(parts, location.CountryCode)
	}

	locationStr := strings.Join(parts, ", ")
	locationStr = strings.TrimSuffix(locationStr, ", ")

	return locationStr
}

func trimURLPrefix(url string) string {
	prefixes := []string{"http://", "https://", "https://www.", "www."}
	for _, prefix := range prefixes {
		if strings.HasPrefix(url, prefix) {
			return strings.TrimPrefix(url, prefix)
		}
	}
	return url
}

func getFirstName(name string) string {
	parts := strings.Split(name, " ")
	if len(parts) > 0 {
		return parts[0]
	}
	return name
}

func getLastName(name string) string {
	parts := strings.Split(name, " ")
	if len(parts) > 1 {
		return strings.Join(parts[1:], " ")
	}
	return name
}

func evaluate(htmlStr string) template.HTML {
	return template.HTML(htmlStr)
}

func lowerEq(s1 string, s2 string) bool {
	return strings.ToLower(s1) == strings.ToLower(s2)
}
