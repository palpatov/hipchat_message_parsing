package fetchers

import (
	"strings"
	"testing"
)

func TestFetchPageTitle(t *testing.T) {
	got, err := FetchPage("http://www.google.com")
	if err != nil || len(got) == 0 || !strings.Contains(got, "<title>Google</title>") {
		t.Error("Fetching Google page title failed")
	}
}
