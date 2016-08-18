package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//
// Index is where we serve the index page
//
func Index(w http.ResponseWriter, _ *http.Request) {
	//todo add instructions to use the service (?Readme.md)
	fmt.Fprintln(w, `Welcome to hipchat_message_parsing service. See below for usage examples

    curl -X POST -H 'Content-Type: application/json' -d "{\"message\": \"Olympics are starting soon; http://www.nbcolympics.com\"}" http://localhost:8080/hipchat_message_parsing

    curl -X POST -H 'Content-Type: application/json' -d "{\"message\": \"Good morning! (megusta) (coffee)\"}" http://localhost:8080/hipchat_message_parsing

    curl -X POST -H 'Content-Type: application/json' -d "{\"message\": \"@bob @john (success) such a cool feature; https://twitter.com/jdorfman/status/430511497475670016\"}" http://localhost:8080/hipchat_message_parsing`)
}

type requestStruct struct {
	Message string
}

//
// MessageParser is where we parse the hipchat message and display the result
//
func MessageParser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t requestStruct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	hipchatMessage := t.Message
	Parse(hipchatMessage, w)
}
