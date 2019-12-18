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

	//subrouter := router.PathPrefix("/api").
	//	HeadersRegexp("Authorization", "Basic [a-zA-Z0-9]{1,128}").
	//	Subrouter()
	//
	//subrouter.HandleFunc("/users/{user_id:[a-zA-Z0-9]{1,64}}/like/{item_id:[a-zA-Z0-9]{1,64}}", users.Like).
	//	Methods(http.MethodPut)
	//subrouter.HandleFunc("/users/{user_id:[a-zA-Z0-9]{1,64}}/like/{item_id:[a-zA-Z0-9]{1,64}}", users.DeleteLike).
	//	Methods(http.MethodDelete)
	//
	//subrouter.HandleFunc("/users/{user_id:[a-zA-Z0-9]{1,64}}/dislike/{item_id:[a-zA-Z0-9]{1,64}}", users.Dislike).
	//	Methods(http.MethodPut)
	//subrouter.HandleFunc("/users/{user_id:[a-zA-Z0-9]{1,64}}/dislike/{item_id:[a-zA-Z0-9]{1,64}}", users.DeleteDislike).
	//	Methods(http.MethodDelete)
	//
	//subrouter.HandleFunc("/users/{user_id:[a-zA-Z0-9]{1,64}}/recommendations", users.Recommendations).
	//	Methods(http.MethodGet)
}

// Handler - check routing and call correct methods
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling url: ", r.URL.Path)

	router.ServeHTTP(w, r)
}
