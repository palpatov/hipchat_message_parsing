package app

import (
	"encoding/json"
	"io"
)

//set these to format our json
const (
	prefix     = ""
	indent     = " "
	escapeHtml = false
)

func FormatJsonToResponse(v interface{}, w io.Writer) {
	en := json.NewEncoder(w)
	en.SetEscapeHTML(escapeHtml)
	en.SetIndent(prefix, indent)
	en.Encode(v)
}
