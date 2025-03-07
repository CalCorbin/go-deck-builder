package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

var decks = make(map[string]Deck)

func GetCards(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	cards, err := FetchCards(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(cards); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateDeck(w http.ResponseWriter, r *http.Request) {
	var deck Deck
	if err := json.NewDecoder(r.Body).Decode(&deck); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deck.ID = strconv.Itoa(rand.Int())
	decks[deck.ID] = deck

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(deck); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deck, ok := decks[id]
	if !ok {
		http.Error(w, "Deck not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(deck); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedDeck Deck
	if err := json.NewDecoder(r.Body).Decode(&updatedDeck); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedDeck.ID = id
	decks[id] = updatedDeck

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(updatedDeck); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := decks[id]; !ok {
		http.Error(w, "Deck not found", http.StatusNotFound)
		return
	}

	delete(decks, id)
	w.WriteHeader(http.StatusNoContent)
}
