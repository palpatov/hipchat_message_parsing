package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//
// This is where we serve the index page
//
func Index(w http.ResponseWriter, r *http.Request) {
	//todo add instructions to use the service (?Readme.md)
	fmt.Fprintln(w, "Welcome!")
}

type request_struct struct {
	Message string
}

//
// This is where we parse the hipchat message and display the result
//
func MessageParser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t request_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	hipchat_message := t.Message
	Parse(hipchat_message, w)
	//fmt.Fprintln(w, "Parse Message:", hipchat_message)
	//fmt.Println(w, Parse(hipchat_message,))
}
