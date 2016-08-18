package parsers

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/palpatov/hipchat_message_parsing/domain"
)

var emoticons_regexp = regexp.MustCompile("\\(([\\w\\d]){1,15}\\)")

/*
parses string for @<user> "mention"
returns "mention" array in json wrapper
*/

func ParseEmoticonsWithFormat(i string) string {
	em := ParseEmoticons(i)
	if em == nil {
		return ""
	}
	//todo factor out json formatting logic
	b, err := json.MarshalIndent(em, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	return string(b)
}

func ParseEmoticons(i string) *domain.EmoticonsType {
	emoticons := emoticons_regexp.FindAllString(i, -1)
	if len(emoticons) == 0 {
		return nil
	}
	for i := 0; i < len(emoticons); i++ {
		emoticons[i] = emoticons[i][1 : len(emoticons[i])-1]
	}
	var ew domain.EmoticonsType
	ew.Emoticons = emoticons
	return &ew
}
