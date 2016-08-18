package parsers

import "github.com/palpatov/hipchat_message_parsing/fetchers"

func WebPageGetTitle(url string) string {
	//todo fill in the remote page title fetching here
	pagesrc := fetchers.FetchPage(url)
	if len(pagesrc) > 0 {
		return ParseTitle(pagesrc)
	}
	return ""
}
