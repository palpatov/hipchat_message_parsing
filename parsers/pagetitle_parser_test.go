package parsers

import "testing"

func TestParseTitle(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"http://www.google.com", "Google"},
		{"", ""},
	} {
		got := WebPageGetTitle(c.in)
		if got != c.want {
			t.Errorf("ParseEmoticons(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
