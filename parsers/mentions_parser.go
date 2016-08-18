package parsers

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/palpatov/hipchat_message_parsing/domain"
)

var mentions_search = regexp.MustCompile("@(\\w+)")

/*
parses string for @<user> "mention"
returns "mention" array in json wrapper
*/

func ParseMentionsWithFormat(i string) string {
	me := ParseMentions(i)
	if me == nil {
		return ""
	}

	b, err := json.MarshalIndent(me, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	return string(b)
}

func ParseMentions(i string) *domain.MentionsType {
	var names []string = mentions_search.FindAllString(i, -1)
	if len(names) == 0 {
		return nil
	}
	for i, v := range names {
		names[i] = v[1:]
	}
	var me domain.MentionsType
	me.MentionNames = names
	return &me
}
