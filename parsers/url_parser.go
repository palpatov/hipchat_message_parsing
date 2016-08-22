package parsers

import (
	"regexp"

	"context"

	"github.com/palpatov/hipchat_message_parsing/domain"
	"github.com/palpatov/hipchat_message_parsing/hcjson"
)

var linksRegexp = regexp.MustCompile("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)")

/*
ParseUrlsWithFormatting parses string for urls
returns URLTupleArray in json wrapper
*/
func ParseUrlsWithFormatting(ctx context.Context, i string) string {
	rw := ParseUrls(ctx, i)
	if rw == nil {
		return ""
	}

	return hcjson.FormatJSONToString(rw)
}

/*
ParseUrls parses string for urls
returns UrlTupleArray
*/
func ParseUrls(ctx context.Context, i string) *domain.URLTupleArray {
	links := linksRegexp.FindAllString(i, -1)
	if len(links) == 0 {
		return nil
	}

	var rw domain.URLTupleArray
	rw.URLs = make([]domain.URLTuple, len(links))
	for i, v := range links {
		var urlt domain.URLTuple
		urlt.Link = v
		urlt.Title, _ = WebPageGetTitle(ctx, v)
		rw.URLs[i] = urlt
	}
	return &rw
}
