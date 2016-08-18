package domain

/*
Mentions is a wrapper for the parsers.ParseMentions
*/
type Mentions struct {
	MentionNames []string `json:"mentions"`
}
