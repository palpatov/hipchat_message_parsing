package fetchers

import (
	"strings"
	"testing"
)

func TestFetchPageTitle(t *testing.T) {
	got := FetchPage("http://www.google.com")
	if len(got) > 0 && strings.Contains(got, "<title>Google<\title>") {
		t.Error("Fetching Google page title failed")
	}
}
