package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicobh15/hb-backend/internal/db/sqlc"
)

type Ingredient struct {
	Item     string `json:"item"`
	Quantity string `json:"quantity"`
}

type Recipe struct {
	Name         string       `json:"name"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions []string     `json:"instructions"`
}
type createRecipeRquest struct {
	Author     string `json:"author" binding:"required"`
	Visibility int32  `json:"visibility" binding:"required"`
	Data       Recipe `json:"data" binding:"required"`
}

type recipeResponse struct {
	RecipeId   pgtype.UUID `json:"recipe_id"`
	Author     string      `json:"author"`
	Visibility int32       `json:"visibility"`
	Data       Recipe      `json:"data"`
}

func (server *Server) createRecipe(ctx *gin.Context) {
	var req createRecipeRquest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	reqData, _ := json.Marshal(req.Data)

	recipe, err := server.store.CreateRecipe(ctx, db.CreateRecipeParams{
		Author:     req.Author,
		Visibility: req.Visibility,
		Data:       reqData,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var recipeData Recipe
	err = json.Unmarshal(recipe.Data, &recipeData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := recipeResponse{
		RecipeId:   recipe.ID,
		Author:     recipe.Author,
		Visibility: recipe.Visibility,
		Data:       recipeData,
	}

	ctx.JSON(http.StatusOK, response)
}

type getRecipeRequest struct {
	RecipeID pgtype.UUID `form:"recipe_id" binding:"required"`
}

func (server *Server) fetchRecipe(ctx *gin.Context) {
	var req getRecipeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	recipe, err := server.store.FetchRecipe(ctx, req.RecipeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var recipeData Recipe
	err = json.Unmarshal(recipe.Data, &recipeData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := recipeResponse{
		RecipeId:   recipe.ID,
		Author:     recipe.Author,
		Visibility: recipe.Visibility,
		Data:       recipeData,
	}

	ctx.JSON(http.StatusOK, response)
}

type deleteRecipeRequest struct {
	RecipeID pgtype.UUID `form:"recipe_id" binding:"required"`
}

func (server *Server) deleteRecipe(ctx *gin.Context) {
	var req deleteRecipeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	recipe, err := server.store.DeleteRecipe(ctx, req.RecipeID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var recipeData Recipe
	err = json.Unmarshal(recipe.Data, &recipeData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := recipeResponse{
		RecipeId:   recipe.ID,
		Author:     recipe.Author,
		Visibility: recipe.Visibility,
		Data:       recipeData,
	}

	ctx.JSON(http.StatusOK, response)
}

type updateRecipeRequest struct {
	RecipeID   pgtype.UUID `json:"recipe_id" binding:"required"`
	Author     string      `json:"author" binding:"required"`
	Visibility int32       `json:"visibility" binding:"required"`
	Data       Recipe      `json:"data" binding:"required"`
}

func (server *Server) updateRecipe(ctx *gin.Context) {
	var req updateRecipeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	reqData, _ := json.Marshal(req.Data)

	recipe, err := server.store.UpdateRecipe(ctx, db.UpdateRecipeParams{
		Author:     req.Author,
		Visibility: req.Visibility,
		Data:       reqData,
		ID:         req.RecipeID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var recipeData Recipe
	err = json.Unmarshal(recipe.Data, &recipeData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := recipeResponse{
		RecipeId:   recipe.ID,
		Author:     recipe.Author,
		Visibility: recipe.Visibility,
		Data:       recipeData,
	}

	ctx.JSON(http.StatusOK, response)
}
