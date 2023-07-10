package app

import (
	"fmt"
	"net/http"
)

// App is the main application
type App struct {
	Server *http.Server
}

// Start starts the application
func (a *App) Start() {
	fmt.Println("Listening on port 8080")
	a.Server.ListenAndServe()
}

// NewApp creates a new application
func NewApp(server *http.Server) *App {
	return &App{Server: server}
}
