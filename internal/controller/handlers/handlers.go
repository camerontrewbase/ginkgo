package handlers

import "github.com/gorilla/mux"

// AddHandlers adds the handlers to the router
func AddHandlers(router *mux.Router) {
	router.HandleFunc("/hello", HelloHandler).Methods("GET")
}
