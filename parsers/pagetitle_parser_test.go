package parsers

import (
	"testing"

	"github.com/palpatov/hipchat_message_parsing/fetchers"
)

func TestCheckLegalURL(t *testing.T) {
	if !fetchers.LegalURL("http://www.google.com") {
		t.Failed()
	}
}

func TestParseTitle(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"http://www.google.com", "Google"},
		{"", ""},
	} {
		got, _ := WebPageGetTitle(c.in)
		if got != c.want {
			t.Errorf("ParseEmoticons(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
