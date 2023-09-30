package main

import (
	"log"
	"net/http"
	"qrcode/handlers"

	"github.com/go-chi/chi"
)

func main() {
	log.Println("Running on port 8087...")

	r := chi.NewRouter()
	r.Post("/generator/", handlers.ViewCodeHandler)
	r.HandleFunc("/", handlers.HomeHandler)

	err := http.ListenAndServe(":8087", r)
	if err != nil {
		log.Panic("Error while starting server.")
	}
}
