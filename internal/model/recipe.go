package model

import (
	"github.com/google/uuid"
)

type Recipe struct {
	ID      uuid.UUID     `json:"id"`
	Author  string        `json:"author"`
	Content RecipeContent `json:"content"`
}

type RecipeContent struct {
	Name         string       `json:"name"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions []string     `json:"instructions"`
	PrepTime     *int         `json:"prepTime,omitempty"`
	TotalTime    *int         `json:"totalTime,omitempty"`
}

type Ingredient struct {
	Item     string `json:"item"`
	Quantity string `json:"quantity"`
}
