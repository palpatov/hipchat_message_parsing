package app

import (
	"io"

	"github.com/palpatov/hipchat_message_parsing/domain"
	"github.com/palpatov/hipchat_message_parsing/hcjson"
	"github.com/palpatov/hipchat_message_parsing/parsers"
)

type ParseResult struct {
	Mentions  []string              `json:"mentions,omitempty"`
	Emoticons []string              `json:"emoticons,omitempty"`
	Links     []domain.UrlTupleType `json:"links,omitempty"`
}

//
// the actual parsing logic goes here
//
func Parse(m string, w io.Writer) {
	var result ParseResult

	me := parsers.ParseMentions(m)
	if me != nil {
		result.Mentions = me.MentionNames
	}
	em := parsers.ParseEmoticons(m)
	if em != nil {
		result.Emoticons = em.Emoticons
	}
	li := parsers.ParseUrls(m)
	if li != nil {
		result.Links = li.URLs
	}

	hcjson.FormatJsonToResponse(result, w)
}
