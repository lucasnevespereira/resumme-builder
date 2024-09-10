package template

import (
	"fmt"
	"html/template"
	"resumme-builder/internal/models"
	"strings"
	"time"

	"github.com/ruang-guru/monday"
)

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

func displayLocationWithSlash(location models.Location) string {
	var parts []string

	if location.City != "" {
		parts = append(parts, location.City)
	}
	if location.CountryCode != "" {
		parts = append(parts, location.CountryCode)
	}
	if location.Region != "" {
		parts = append(parts, location.Region)
	}

	locationStr := strings.Join(parts, "/ ")
	locationStr = strings.TrimSuffix(locationStr, "/ ")

	return locationStr
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

var formats = []string{
	"2006-01-02",
	"2006-01",
	"January 2 2006",
	"January 2006",
	"2006",
}

func formatDate(layout string, date string, locale string) string {
	for _, format := range formats {
		t, err := time.Parse(format, date)
		if err == nil {
			return monday.Format(t, layout, monday.Locale(locale))
		}
	}

	panic(fmt.Sprintf("Source string date format could not be recognized, valid formats are: %v", strings.Join(formats, ", ")))
}
