// Package ui embeds the themes and translations in the binary.
package ui

import "embed"

// all: keeps the _<theme>.gohtml root files.
//
//go:embed all:templates locales
var FS embed.FS
