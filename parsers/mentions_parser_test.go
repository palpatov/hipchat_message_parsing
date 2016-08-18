package parsers

import "testing"

func TestParseMentions(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"@chris you around?", `{
  "mentions": [
    "chris"
  ]
}`},
		{"@Chris are you up for lunch with @Bob?", `{
  "mentions": [
    "Chris",
    "Bob"
  ]
}`},
		{"", ""},
	} {
		got := ParseMentionsWithFormat(c.in)
		if got != c.want {
			t.Errorf("ParseMentions(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
