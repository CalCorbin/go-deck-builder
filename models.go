package main

type Card struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Price string `json:"price"`
}

type Deck struct {
	ID    string `json:"ID"`
	Name  string `json:"name"`
	Cards []Card `json:"cards"`
}
