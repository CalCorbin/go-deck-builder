package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/cards", GetCards).Methods(http.MethodGet)
	router.HandleFunc("/api/decks", CreateDeck).Methods(http.MethodPost)
	router.HandleFunc("/api/decks/{id}", GetDeck).Methods(http.MethodGet)
	router.HandleFunc("/api/decks/{id}", UpdateDeck).Methods(http.MethodPut)
	router.HandleFunc("/api/decks/{id}", DeleteDeck).Methods(http.MethodDelete)

	return router
}
