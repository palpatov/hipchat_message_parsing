//
// json formatting settings go here
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
	EscapeHtml = false
)

func FormatJsonToResponse(v interface{}, w io.Writer) {
	en := json.NewEncoder(w)
	en.SetEscapeHTML(EscapeHtml)
	en.SetIndent(Prefix, Indent)
	en.Encode(v)
}

func FormatJsonToString(v interface{}) string {
	var b bytes.Buffer
	FormatJsonToResponse(v, &b)
	return b.String()
}
