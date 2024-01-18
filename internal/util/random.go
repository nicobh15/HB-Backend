package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Ingredient struct {
	Item     string `json:"item"`
	Quantity string `json:"quantity"`
}

type RecipeData struct {
	Name         string       `json:"name"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions []string     `json:"instructions"`
}

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

func RandomInt(min, max int64) int64 {
	return (min + rand.Int63n(max-min+1))
}
func RandomInt32(min, max int32) int32 {
	return (min + rand.Int31n(max-min+1))
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail() string {
	return (RandomString(10) + "@test.com")
}

func RandomUserName() string {
	return (RandomString(10))
}

func RandomName() string {
	return (RandomString(7))
}

func RandomAddress() pgtype.Text {
	return (pgtype.Text{String: RandomString(10), Valid: true})
}

func RandomUUID() pgtype.UUID {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
	}

	var pgUUID pgtype.UUID
	copy(pgUUID.Bytes[:], newUUID[:])
	pgUUID.Valid = true

	return (pgUUID)
}

func RandomCategory() string {
	return (RandomString(10))
}

func RandomLocation() pgtype.Text {
	return (pgtype.Text{String: RandomString(10), Valid: true})
}

func RandomDate() pgtype.Date {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)

	// Calculate time range in seconds
	delta := end.Unix() - start.Unix()

	// Generate a random number of seconds since start
	sec := rand.Int63n(delta) + start.Unix()

	// Convert to time.Time
	randomTime := time.Unix(sec, 0)

	return (pgtype.Date{Time: randomTime, Valid: true})
}

func RandomIngredientQuantity() string {
	unit := RandomString(5)
	quantity := RandomInt(1, 100)

	finalQuantity := unit + " " + strconv.FormatInt(quantity, 10)
	return (finalQuantity)
}

func RandomRecipeData() []byte {

	name := RandomName()

	ingredientCount := RandomInt32(1, 10)
	instructionCount := RandomInt32(1, 10)

	ingredients := make([]Ingredient, ingredientCount)
	instructions := make([]string, instructionCount)

	for i := range ingredients {
		ingredients[i].Item = RandomName()
		ingredients[i].Quantity = RandomIngredientQuantity()
	}

	for i := range instructions {
		instructions[i] = RandomString(100)
	}

	recipe := RecipeData{
		Name:         name,
		Ingredients:  ingredients,
		Instructions: instructions,
	}

	recipeJSON, err := json.Marshal(recipe)
	if err != nil {
		fmt.Println("Error marshaling recipe to JSON:", err)
		return nil
	}

	return recipeJSON
}
