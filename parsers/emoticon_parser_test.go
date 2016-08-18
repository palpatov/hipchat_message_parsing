package parsers

import "testing"

func TestParseEmoticons(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"Good morning! (megusta) (coffee11)", `{
  "emoticons": [
    "megusta",
    "coffee11"
  ]
}`},
		{"", ""},
		{"Good morning! (coffeeeeeeeeeeee)", ""},
	} {
		got := ParseEmoticonsWithFormat(c.in)
		if got != c.want {
			t.Errorf("ParseEmoticons(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
