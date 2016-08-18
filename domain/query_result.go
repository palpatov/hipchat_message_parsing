package domain

//
// what the resultset of the query looks like
//
type ParseResult struct {
	Mentions  []string       `json:"mentions,omitempty"`
	Emoticons []string       `json:"emoticons,omitempty"`
	Links     []UrlTupleType `json:"links,omitempty"`
}
