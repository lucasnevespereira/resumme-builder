package template

import (
	"html/template"
	"resumme-builder/internal/models"
	"strings"
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
	return strings.EqualFold(strings.ToLower(s1), strings.ToLower(s2))
}
