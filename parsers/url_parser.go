package parsers

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/palpatov/hipchat_message_parsing/domain"
)

var links_regexp = regexp.MustCompile("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)")

/*
parses string for @<user> "mention"
returns "mention" array in json wrapper
*/
func ParseUrlsWithFormatting(i string) string {
	rw := ParseUrls(i)
	if rw == nil {
		return ""
	}

	//todo factor out json formatting logic
	b, err := json.MarshalIndent(rw, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	return string(b)
}

func ParseUrls(i string) *domain.UrlTupleArrayType {
	links := links_regexp.FindAllString(i, -1)
	if len(links) == 0 {
		return nil
	}

	var rw domain.UrlTupleArrayType
	rw.URLs = make([]domain.UrlTupleType, len(links))
	for i, v := range links {
		var urlt domain.UrlTupleType
		urlt.Link, urlt.Title = v, WebPageGetTitle(v)
		rw.URLs[i] = urlt
	}
	return &rw
}
