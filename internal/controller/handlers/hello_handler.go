package handlers

import (
	"encoding/json"
	model "github.com/camerontrewbase/ginkgo/internal/domain/hello"
	"net/http"
)

// HelloHandler handles the hello request
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{Message: "Hello World"})
}
