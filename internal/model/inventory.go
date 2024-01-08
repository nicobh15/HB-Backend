package model

import (
	"time"

	"github.com/google/uuid"
)

// InventoryItem represents an item in the inventory.
type InventoryItem struct {
	ItemID         uuid.UUID  `json:"itemID"`
	HouseholdID    uuid.UUID  `json:"householdID"`
	Category       string     `json:"category"` // e.g., "Kitchen", "Bar", "Storage"
	Name           string     `json:"name"`
	Quantity       int        `json:"quantity"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	PurchaseDate   time.Time  `json:"purchaseDate"`
}
