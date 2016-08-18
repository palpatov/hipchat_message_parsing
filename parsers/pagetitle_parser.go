package parsers

import "regexp"

var titleRegexp = regexp.MustCompile("<title>(.+)<\\/title>")

/*
ParseTitle parses title string from the stringified pagesource
as per the regex above
*/
func ParseTitle(src string) string {
	title := titleRegexp.FindString(src)
	if len(title) > 0 {
		title = title[len("<title>") : len(title)-len("<\\title>")]
	}
	return title
}
