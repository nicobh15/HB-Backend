package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicobh15/HomeBuddy-Backend/internal/db/sqlc"
)

type CreateHouseholdRequest struct {
	HouseholdName string      `json:"household_name" binding:"required"`
	Address       pgtype.Text `json:"address" `
}

func (server *Server) createHousehold(ctx *gin.Context) {
	var req CreateHouseholdRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	household, err := server.store.CreateHousehold(ctx, db.CreateHouseholdParams{
		HouseholdName: req.HouseholdName,
		Address:       req.Address,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, household)
}

type GetHouseholdRequest struct {
	HouseholdID pgtype.UUID `form:"household_id" binding:"required"`
}

func (server *Server) fetchHousehold(ctx *gin.Context) {
	var req GetHouseholdRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	household, err := server.store.FetchHousehold(ctx, req.HouseholdID)

	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, household)
}

type ListHouseholdsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *Server) listHouseholds(ctx *gin.Context) {
	var req ListHouseholdsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	households, err := server.store.ListHouseholds(ctx, db.ListHouseholdsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, households)
}

type DeleteHouseholdRequest struct {
	HouseholdID pgtype.UUID `form:"household_id" binding:"required"`
}

func (server *Server) deleteHousehold(ctx *gin.Context) {
	var req DeleteHouseholdRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	household, err := server.store.DeleteHousehold(ctx, req.HouseholdID)

	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, household)
}

type UpdateHouseholdRequest struct {
	HouseholdID   pgtype.UUID `json:"household_id" binding:"required"`
	HouseholdName string      `json:"household_name"`
	Address       pgtype.Text `json:"address" `
}

func (server *Server) updateHousehold(ctx *gin.Context) {
	var req UpdateHouseholdRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	household, err := server.store.UpdateHousehold(ctx, db.UpdateHouseholdParams{
		HouseholdName: req.HouseholdName,
		Address:       req.Address,
		HouseholdID:   req.HouseholdID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, household)
}
