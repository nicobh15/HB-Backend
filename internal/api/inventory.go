package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicobh15/hb-backend/internal/db/sqlc"
	"github.com/nicobh15/hb-backend/internal/token"
)

type createInventoryItemRequest struct {
	HouseholdID pgtype.UUID `json:"household_id" binding:"required"`
	ItemName    string      `json:"item_name" binding:"required"`
	Quantity    int32       `json:"quantity" binding:"required"`
	Category    string      `json:"category" binding:"required"`
	Location    pgtype.Text `json:"location" binding:"required"`
}

func (server *Server) createInventoryItem(ctx *gin.Context) {
	var req createInventoryItemRequest
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
	inventoryItem, err := server.store.CreateInventoryItem(ctx, db.CreateInventoryItemParams{
		HouseholdID: req.HouseholdID,
		Category:    req.Category,
		Name:        req.ItemName,
		Quantity:    req.Quantity,
		Location:    req.Location,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, inventoryItem)
}

type getHouseholdInventoryRequest struct {
	HouseholdID pgtype.UUID `form:"household_id" binding:"required"`
	Limit       int32       `form:"limit,default=10" binding:"max=100"`
	Offset      int32       `form:"offset,default=0" binding:"max=100"`
}

func (server *Server) fetchHouseholdInventory(ctx *gin.Context) {
	var req getHouseholdInventoryRequest
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
	inventory, err := server.store.ListInventoryItems(ctx, db.ListInventoryItemsParams{
		HouseholdID: req.HouseholdID,
		Limit:       req.Limit,
		Offset:      req.Offset,
	})

	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, inventory)
}

type deleteInventoryItemRequest struct {
	ItemID      pgtype.UUID `form:"item_id" binding:"required"`
	HouseholdID pgtype.UUID `form:"household_id" binding:"required"`
}

func (server *Server) deleteInventoryItem(ctx *gin.Context) {
	var req deleteInventoryItemRequest
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

	inventoryItem, err := server.store.DeleteInventoryItem(ctx, req.ItemID)

	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, inventoryItem)
}

type updateInventoryItemRequest struct {
	ItemID       pgtype.UUID `json:"item_id" binding:"required"`
	HouseholdID  pgtype.UUID `json:"household_id" binding:"required"`
	ItemName     string      `json:"item_name" binding:"required"`
	Quantity     int32       `json:"quantity" binding:"required"`
	Category     string      `json:"category" binding:"required"`
	Location     pgtype.Text `json:"location" binding:"required"`
	Expiration   pgtype.Date `json:"expiration_date" binding:"required"`
	PurchaseDate pgtype.Date `json:"purchase_date" binding:"required"`
}

func (server *Server) updateInventoryItem(ctx *gin.Context) {
	var req updateInventoryItemRequest
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

	inventoryItem, err := server.store.UpdateInventoryItem(ctx, db.UpdateInventoryItemParams{
		ItemID:         req.ItemID,
		HouseholdID:    req.HouseholdID,
		Name:           req.ItemName,
		Quantity:       req.Quantity,
		Category:       req.Category,
		Location:       req.Location,
		ExpirationDate: req.Expiration,
		PurchaseDate:   req.PurchaseDate,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, inventoryItem)
}
