package parsers

import (
	"context"

	"github.com/palpatov/hipchat_message_parsing/fetchers"
)

/*
WebPageGetTitle retrieves the title of the web resource by retrieving the resource first.
returns title or title and an error if unsuccessfull
*/
func WebPageGetTitle(ctx context.Context, url string) (string, error) {

	//todo what do we do for long running fetches?
	//we need some non blocking way of doing this
	pagesrc, err := fetchers.FetchPage(ctx, url)
	if len(pagesrc) > 0 && err == nil {
		return ParseTitle(pagesrc), nil
	}
	return "", err
}
