package main

import (
	"net/http"
	"qrcode/handlers"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()
	r.Post("/generator/", handlers.ViewCodeHandler)
	r.HandleFunc("/", handlers.HomeHandler)

	http.ListenAndServe(":8088", r)
}
