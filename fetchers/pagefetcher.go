package fetchers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"regexp"
)

var links_regexp = regexp.MustCompile("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)")

func LegalUrl(url string) bool {
	return links_regexp.MatchString(url)
}

func FetchPage(url string) string {
	// check url for legality
	if !LegalUrl(url) {
		return ""
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTML:\n\n", string(err.Error()))
		return ""
	}

	bytes, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println("HTML:\n\n", string(bytes))

	resp.Body.Close()

	return string(bytes)
}
