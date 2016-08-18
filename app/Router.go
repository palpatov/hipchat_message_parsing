package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

/*
NewRouter creates web application router
*/
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, r := range myroutes {
		router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(r.HandlerFunc)
	}

	return router
}

var myroutes = routes{
	route{
		"Index",
		"GET",
		"/",
		Index,
	},
	route{
		"MessageParser",
		"POST",
		"/hipchat_message_parsing",
		MessageParser,
	},
}
