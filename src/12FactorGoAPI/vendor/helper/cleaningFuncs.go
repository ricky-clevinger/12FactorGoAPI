package helper

import (
	"html"
	"strings"
)

func HtmlClean(toClean string) string {

	toClean = strings.TrimSpace(toClean)
	cleanedString := html.EscapeString(toClean)

	return cleanedString
}
