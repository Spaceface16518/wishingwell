package router

import (
	"github.com/gorilla/mux"
	"wishingwell/www/wishlist"
)

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/wishlist", wishlist.Handler)

	return
}
