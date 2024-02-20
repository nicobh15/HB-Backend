package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicobh15/HomeBuddy-Backend/internal/db/sqlc"
	"github.com/nicobh15/HomeBuddy-Backend/internal/token"
)

type createHouseholdRequest struct {
	HouseholdName string      `json:"household_name" binding:"required"`
	Address       pgtype.Text `json:"address" `
}

func (server *Server) createHousehold(ctx *gin.Context) {
	var req createHouseholdRequest
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

type getHouseholdRequest struct {
	HouseholdID pgtype.UUID `form:"household_id" binding:"required"`
}

func (server *Server) fetchHousehold(ctx *gin.Context) {
	var req getHouseholdRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.User.HouseholdID != req.HouseholdID {
		err := fmt.Errorf("unauthorized Access")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
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

type listHouseholdsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *Server) listHouseholds(ctx *gin.Context) {
	var req listHouseholdsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.User.Role != "admin" {
		err := fmt.Errorf("unauthorized Access")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
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

type deleteHouseholdRequest struct {
	HouseholdID pgtype.UUID `form:"household_id" binding:"required"`
}

func (server *Server) deleteHousehold(ctx *gin.Context) {
	var req deleteHouseholdRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.User.HouseholdID != req.HouseholdID {
		err := fmt.Errorf("unauthorized Access")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
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

type updateHouseholdRequest struct {
	HouseholdID   pgtype.UUID `json:"household_id" binding:"required"`
	HouseholdName string      `json:"household_name" binding:"required"`
	Address       pgtype.Text `json:"address" binding:"required"`
}

func (server *Server) updateHousehold(ctx *gin.Context) {
	var req updateHouseholdRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.User.HouseholdID != req.HouseholdID {
		err := fmt.Errorf("unauthorized Access")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
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
