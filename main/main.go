package main

import (
	"log"
	"net/http"
	"wishingwell/internal/router"
)

func main() {
	r := router.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
