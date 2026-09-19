package render

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"path"
	"strings"
)

const defaultLang = "en"

// loadLabels reads every locales/<lang>.json into lang -> key -> text.
func loadLabels(ui fs.FS) (map[string]map[string]string, error) {
	files, err := fs.Glob(ui, "locales/*.json")
	if err != nil {
		return nil, err
	}

	labels := map[string]map[string]string{}
	for _, file := range files {
		data, err := fs.ReadFile(ui, file)
		if err != nil {
			return nil, err
		}
		var l map[string]string
		if err := json.Unmarshal(data, &l); err != nil {
			return nil, fmt.Errorf("%s: %w", file, err)
		}
		labels[strings.TrimSuffix(path.Base(file), ".json")] = l
	}

	if _, ok := labels[defaultLang]; !ok {
		return nil, errors.New("missing locales/" + defaultLang + ".json")
	}
	return labels, nil
}

// labelsFor falls back to English for a language without a locale file.
func (r *Renderer) labelsFor(lang string) map[string]string {
	if l, ok := r.labels[lang]; ok {
		return l
	}
	if l, ok := r.labels[baseLang(lang)]; ok {
		return l
	}
	return r.labels[defaultLang]
}

// baseLang turns regional codes like "fr_FR" or "fr-FR" into "fr".
// Labels and dates both go through it, so they always agree.
func baseLang(lang string) string {
	base, _, _ := strings.Cut(strings.ReplaceAll(lang, "-", "_"), "_")
	return strings.ToLower(base)
}
