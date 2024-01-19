package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicobh15/HomeBuddy-Backend/internal/db/sqlc"
)

type CreateUserRequest struct {
	Username     string      `json:"username" binding:"required"`
	Email        string      `json:"email" binding:"required"`
	FirstName    string      `json:"first_name" binding:"required"`
	PasswordHash string      `json:"password_hash" binding:"required"`
	Role         string      `json:"role" binding:"required"`
	HouseholdID  pgtype.UUID `json:"household_id"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.store.CreateUser(ctx, db.CreateUserParams{
		Username:     req.Username,
		Email:        req.Email,
		FirstName:    req.FirstName,
		PasswordHash: req.PasswordHash,
		Role:         req.Role,
		HouseholdID:  req.HouseholdID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}
