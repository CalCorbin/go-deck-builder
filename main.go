package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	handler := cors.Default().Handler(router)
	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
