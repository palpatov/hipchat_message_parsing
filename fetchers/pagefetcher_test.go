package fetchers

import (
	"context"
	"strings"
	"testing"
)

func TestFetchPageTitle(t *testing.T) {
	ctx := context.Background()
	got, err := FetchPage(ctx, "http://www.google.com")
	if err != nil || len(got) == 0 || !strings.Contains(got, "<title>Google</title>") {
		t.Error("Fetching Google page title failed")
	}
}
