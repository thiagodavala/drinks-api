package models

type Cocktail struct {
	ID          int    `json:"Id"`
	Name        string `json:"Name"`
	Ingredients string `json:"Ingredients"`
	Method      string `json:"Method"`
	Garnish     string `json:"Garnish"`
}
