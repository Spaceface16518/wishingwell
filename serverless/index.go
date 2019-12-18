package serverless

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	wishrouter "wishingwell/internal/router"
)

var router *mux.Router

func init() {
	router = wishrouter.NewRouter()
}

// Handler - check routing and call correct methods
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling url: ", r.URL.Path)

	router.ServeHTTP(w, r)
}
