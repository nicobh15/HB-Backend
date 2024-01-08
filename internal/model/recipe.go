package model

import (
	"github.com/google/uuid"
)

// Recipe represents a recipe in the system.
type Recipe struct {
	ID      uuid.UUID     `json:"id"`
	Author  string        `json:"author"`
	Content RecipeContent `json:"content"`
}

type RecipeContent struct {
	Name         string       `json:"name"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions []string     `json:"instructions"`
}

type Ingredient struct {
	Item     string `json:"item"`
	Quantity string `json:"quantity"`
}
