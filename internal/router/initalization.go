package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"wishingwell/internal/templates"
	"wishingwell/www/wishlist"
)

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/wishlist", wishlist.Handler)
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		templates.Templates.ExecuteTemplate(writer, "index.gohtml", nil)
	})

	return
}
