//
// Package hcjson wraps formatting settings
// and helper functions
//
package hcjson

import (
	"bytes"
	"encoding/json"
	"io"
)

//set these to format our json
const (
	Prefix     = ""
	Indent     = "  "
	EscapeHTML = false
)

/*
FormatJSONToResponse formats object to a stream with
JSON formating specified via constants above
*/
func FormatJSONToResponse(v interface{}, w io.Writer) {
	en := json.NewEncoder(w)
	en.SetEscapeHTML(EscapeHTML)
	en.SetIndent(Prefix, Indent)
	en.Encode(v)
}

/*
FormatJSONToString formats object to a string of
JSON format specified via constants above
*/
func FormatJSONToString(v interface{}) string {
	var b bytes.Buffer
	FormatJSONToResponse(v, &b)
	return b.String()
}
