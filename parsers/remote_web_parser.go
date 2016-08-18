package parsers

import "github.com/palpatov/hipchat_message_parsing/fetchers"

func WebPageGetTitle(url string) string {

	//todo what do we do for long running fetches?
	//we need some non blocking way of doing this
	pagesrc := fetchers.FetchPage(url)
	if len(pagesrc) > 0 {
		return ParseTitle(pagesrc)
	}
	return ""
}
