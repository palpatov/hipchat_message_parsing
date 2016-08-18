package parsers

import "testing"

func TestParseLinks(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"Olympics are starting soon; http://www.nbcolympics.com", `{
  "links": [
    {
      "url": "http://www.nbcolympics.com",
      "title": "2016 Rio Olympic Games | NBC Olympics"
    }
  ]
}`},
		{"", ""},
	} {
		got := ParseUrlsWithFormatting(c.in)
		if got != c.want {
			t.Errorf("ParseEmoticons(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
