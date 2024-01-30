package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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

type GetUserRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (server *Server) fetchUserByEmail(ctx *gin.Context) {
	var req GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.store.FetchUserByUserName(ctx, req.Username)

	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type ListUsersRequest struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req ListUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	users, err := server.store.ListUsers(ctx, db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, users)
}

type ListUsersByHouseholdRequest struct {
	HouseholdID pgtype.UUID `form:"household_id" binding:"required"`
	PageId      int32       `form:"page_id" binding:"required,min=1"`
	PageSize    int32       `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *Server) listUsersByHousehold(ctx *gin.Context) {
	var req ListUsersByHouseholdRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	users, err := server.store.ListHouseholdMembers(ctx, db.ListHouseholdMembersParams{
		HouseholdID: req.HouseholdID,
		Limit:       req.PageSize,
		Offset:      (req.PageId - 1) * req.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, users)

}
