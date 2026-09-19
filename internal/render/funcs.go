package render

import (
	"html/template"
	"resumme-builder/internal/models"
	"strings"
	"time"

	"github.com/ruang-guru/monday"
)

var funcs = template.FuncMap{
	"isLast":                            isLast,
	"displayLocation":                   displayLocation,
	"displayLocationWithSlash":          displayLocationWithSlash,
	"displayLocationWithHyphen":         displayLocationWithHyphen,
	"displayLocationWithCommaAndHyphen": displayLocationWithCommaAndHyphen,
	"displayLocationWithCommaAndSpace":  displayLocationWithCommaAndSpace,
	"displayLocationWithSpaces":         displayLocationWithSpaces,
	"trimURLPrefix":                     trimURLPrefix,
	"getFirstName":                      getFirstName,
	"getLastName":                       getLastName,
	"evaluate":                          evaluate,
	"lowerEq":                           lowerEq,
	"lower":                             lower,
	"formatDate":                        formatDate,
	"paragraphLineFeeds":                paragraphLineFeeds,
	"imageSrc":                          imageSrc,
}

func isLast(index, length int) bool {
	return index == length-1
}

func formatLocation(location models.Location, separatorRegion, separatorCountry string) string {
	var result string

	if location.City != "" {
		result = location.City
	}
	if location.Region != "" {
		if result != "" {
			result += separatorRegion
		}
		result += location.Region
	}
	if location.CountryCode != "" {
		if result != "" {
			result += separatorCountry
		}
		result += location.CountryCode
	}

	return result
}

func displayLocation(location models.Location) string {
	return formatLocation(location, ", ", ", ")
}

func displayLocationWithSlash(location models.Location) string {
	return formatLocation(location, "/ ", "/ ")
}

func displayLocationWithHyphen(location models.Location) string {
	return formatLocation(location, " - ", " - ")
}

func displayLocationWithCommaAndHyphen(location models.Location) string {
	return formatLocation(location, ", ", " - ")
}

func displayLocationWithCommaAndSpace(location models.Location) string {
	return formatLocation(location, ", ", " ")
}

func displayLocationWithSpaces(location models.Location) string {
	return formatLocation(location, " ", " ")
}

func trimURLPrefix(url string) string {
	prefixes := []string{"http://", "https://", "https://www.", "www.", ""}
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
	return strings.EqualFold(strings.ToLower(s1), strings.ToLower(s2))
}

func lower(s string) string {
	return strings.ToLower(s)
}

const yearOnlyFormat = "2006"

var formats = []string{
	"2006-01-02",
	"2006-01",
	"January 2 2006",
	"January 2006",
	yearOnlyFormat,
}

// monday expects full locales ("fr_FR"), while resume data carries short codes ("fr").
var mondayLocales = map[string]monday.Locale{
	"fr": monday.LocaleFrFR,
	"en": monday.LocaleEnUS,
}

func formatDate(layout string, date string, locale string) string {
	mondayLocale, ok := mondayLocales[locale]
	if !ok {
		mondayLocale = monday.LocaleEnUS
	}

	for _, format := range formats {
		t, err := time.Parse(format, date)
		if err != nil {
			continue
		}
		// A source reduced to a year is returned as is: formatting it would
		// attribute a January that the data never carried.
		if format == yearOnlyFormat {
			return date
		}
		return monday.Format(t, layout, mondayLocale)
	}
	return date
}

// html/template rewrites data: URIs to "#ZgotmplZ", which breaks embedded images.
func imageSrc(src string) template.URL {
	return template.URL(src)
}

func paragraphLineFeeds(text string) template.HTML {
	output := strings.ReplaceAll(text, "\n", "</p>\n<p>")
	return template.HTML(output)
}
