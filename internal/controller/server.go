package controller

import (
	"github.com/camerontrewbase/ginkgo/internal/controller/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

// NewServer creates a new server
func NewServer() *http.Server {
	router := mux.NewRouter()
	handlers.AddHandlers(router)
	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}
