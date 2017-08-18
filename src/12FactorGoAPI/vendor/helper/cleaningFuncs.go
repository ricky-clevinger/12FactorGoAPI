package helper

import (
	"html"
)

func HtmlClean(toClean string) string {

	cleanedString := html.EscapeString(toClean)

	return cleanedString
}
