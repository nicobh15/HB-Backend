package model

import (
	"time"

	"github.com/google/uuid"
)

type InventoryItem struct {
	ItemID         uuid.UUID  `json:"itemID"`
	HouseholdID    uuid.UUID  `json:"householdID"`
	Category       string     `json:"category"`
	Name           string     `json:"name"`
	Quantity       int        `json:"quantity"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	PurchaseDate   time.Time  `json:"purchaseDate"`
}
