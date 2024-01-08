package model

import (
	"github.com/google/uuid"
)

type Household struct {
	HouseholdID uuid.UUID `json:"householdID"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
}
