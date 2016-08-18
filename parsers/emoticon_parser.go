package parsers

import (
	"regexp"

	"github.com/palpatov/hipchat_message_parsing/domain"
	"github.com/palpatov/hipchat_message_parsing/hcjson"
)

var emoticonsRegexp = regexp.MustCompile("\\(([\\w\\d]){1,15}\\)")

//
// ParseEmoticonsWithFormat parses emoticons from imput string and returns
// json representaiton, as
//
func ParseEmoticonsWithFormat(i string) string {
	em := ParseEmoticons(i)
	if em == nil {
		return ""
	}

	return hcjson.FormatJSONToString(em)
}

//
// ParseEmoticons collects emoticons from a string
// returns address of EmoticonsBag struct (see domain)
// or nil if none detected
//
func ParseEmoticons(i string) *domain.EmoticonsBag {
	emoticons := emoticonsRegexp.FindAllString(i, -1)
	if len(emoticons) == 0 {
		return nil
	}
	for i := 0; i < len(emoticons); i++ {
		emoticons[i] = emoticons[i][1 : len(emoticons[i])-1]
	}
	var ew domain.EmoticonsBag
	ew.Emoticons = emoticons
	return &ew
}
