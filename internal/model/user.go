package model

import "github.com/google/uuid"

// User represents a user in the system.
type User struct {
	UserID       uuid.UUID `json:"userID"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"passwordHash"`
	Role         string    `json:"role"`
	HouseholdID  uuid.UUID `json:"householdID"`
}
