package fetchers

import (
	"testing"

	"github.com/palpatov/hipchat_message_parsing/parsers"
)

func TestFetchPageTitle(t *testing.T) {
	got := FetchPage("http://www.google.com")
	if len(got) > 0 && parsers.ParseTitle(got) == "Google" {
		t.Error("Fetching Google page title failed")
	}
}
