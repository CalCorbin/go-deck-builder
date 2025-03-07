package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const scryfallApi = "https://api.scryfall.com/cards/search?q="

func FetchCards(query string) ([]Card, error) {
	// Hit the scryfall API and create our response and error variables
	resp, err := http.Get(scryfallAPI + query)

	// Return early for an error
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Return early again if we get don't get a good response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch cards: %s", resp.Status)
	}

	// Create our result data with response
	var result struct {
		Data []Card `json:"data"`
	}

	// Decode json response and handle any errors
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	// Return scryfall data
	return result.Data, nil
}
