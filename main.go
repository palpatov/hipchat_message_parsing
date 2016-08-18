package main

import (
	"log"
	"net/http"

	"github.com/palpatov/hipchat_message_parsing/app"
)

func main() {

	router := app.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

//func main() {
//	var ret string = parsers.ParseUrls("Olympics are starting soon; http://www.nbcolympics.com")
//	fmt.Println(ret)
//}
