package parsers

import (
	"regexp"

	"github.com/palpatov/hipchat_message_parsing/domain"
	"github.com/palpatov/hipchat_message_parsing/hcjson"
)

var mentionsSearchRegexp = regexp.MustCompile("@(\\w+)")

/*
ParseMentionsWithFormat parses string for @<user> "mention"
returns "mention" array in json wrapper or empty string if none detected
*/
func ParseMentionsWithFormat(i string) string {
	me := ParseMentions(i)
	if me == nil {
		return ""
	}

	return hcjson.FormatJSONToString(me)
}

/*
ParseMentions parses string for @<user> "mention"
returns MentionsType struct or nil if none detected
*/
func ParseMentions(i string) *domain.Mentions {
	names := mentionsSearchRegexp.FindAllString(i, -1)
	if len(names) == 0 {
		return nil
	}
	for i, v := range names {
		names[i] = v[1:]
	}
	var me domain.Mentions
	me.MentionNames = names
	return &me
}
