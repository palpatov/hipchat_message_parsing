package parsers

import "regexp"

var title_regexp = regexp.MustCompile("<title>(.+)<\\/title>")

func ParseTitle(src string) string {
	title := title_regexp.FindString(src)
	if len(title) > 0 {
		title = title[len("<title>") : len(title)-len("<\\title>")]
	}
	return title
}
