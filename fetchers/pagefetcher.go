package fetchers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"errors"
	"regexp"
)

var linksRegexp = regexp.MustCompile("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)")

//
// LegalURL checks for correct url formatting
//
func LegalURL(url string) bool {
	return linksRegexp.MatchString(url)
}

//
// FetchPage fetches the contents of a webpage
// returns the contents in a string or
// an empty string and error if page had
// failed to retrieve
//
func FetchPage(url string) (string, error) {
	// check url for legality
	if !LegalURL(url) {
		return "", errors.New("Illegal URL")
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTML:", string(err.Error()))
		return "", err
	}

	bytes, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	return string(bytes), nil
}
