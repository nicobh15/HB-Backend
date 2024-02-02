package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicobh15/HomeBuddy-Backend/internal/db/sqlc"
	"github.com/nicobh15/HomeBuddy-Backend/internal/util"
)

type createUserRequest struct {
	Username    string      `json:"username" binding:"required"`
	Email       string      `json:"email" binding:"required"`
	FirstName   string      `json:"first_name" binding:"required"`
	Password    string      `json:"password" binding:"required,min=10"`
	Role        string      `json:"role" binding:"required"`
	HouseholdID pgtype.UUID `json:"household_id"`
}

type createUserResponse struct {
	Username    string             `json:"username"`
	Email       string             `json:"email"`
	FirstName   string             `json:"first_name"`
	Role        string             `json:"role"`
	HouseholdID pgtype.UUID        `json:"household_id"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.store.CreateUser(ctx, db.CreateUserParams{
		Username:     req.Username,
		Email:        req.Email,
		FirstName:    req.FirstName,
		PasswordHash: hashedPassword,
		Role:         req.Role,
		HouseholdID:  req.HouseholdID,
	})

	rsp := createUserResponse{
		Username:    user.Username,
		Email:       user.Email,
		FirstName:   user.FirstName,
		Role:        user.Role,
		HouseholdID: user.HouseholdID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp)
}

type getUserRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (server *Server) fetchUserByEmail(ctx *gin.Context) {
	var req getUserRequest
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

type listUsersRequest struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req listUsersRequest
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

type listUsersByHouseholdRequest struct {
	HouseholdID pgtype.UUID `form:"household_id" binding:"required"`
	PageId      int32       `form:"page_id" binding:"required,min=1"`
	PageSize    int32       `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *Server) listUsersByHousehold(ctx *gin.Context) {
	var req listUsersByHouseholdRequest
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

type deleteUserRequest struct {
	Email string `uri:"email" binding:"required"`
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.store.DeleteUser(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type updateUserRequest struct {
	Username    string      `json:"username" binding:"required"`
	Email       string      `json:"email" binding:"required,email"`
	FirstName   string      `json:"first_name" binding:"required"`
	Password    string      `json:"password" binding:"required,min=8"`
	Role        string      `json:"role" binding:"required"`
	HouseholdID pgtype.UUID `json:"household_id" binding:"required"`
	UserID      pgtype.UUID `json:"user_id" binding:"required"`
}
type updateUserResponse struct {
	Username    string             `json:"username"`
	Email       string             `json:"email"`
	FirstName   string             `json:"first_name"`
	Role        string             `json:"role"`
	HouseholdID pgtype.UUID        `json:"household_id"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.store.UpdateUser(ctx, db.UpdateUserParams{
		Username:     req.Username,
		Email:        req.Email,
		FirstName:    req.FirstName,
		PasswordHash: hashedPassword,
		Role:         req.Role,
		HouseholdID:  req.HouseholdID,
		UserID:       req.UserID,
	})
	rsp := updateUserResponse{
		Username:    user.Username,
		Email:       user.Email,
		FirstName:   user.FirstName,
		Role:        user.Role,
		HouseholdID: user.HouseholdID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp)
}
