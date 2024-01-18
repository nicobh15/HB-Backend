package db

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/nicobh15/HomeBuddy-Backend/internal/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomRecipe(t *testing.T) Recipe {
	author := CreateRandomUser(t)

	args := CreateRecipeParams{
		AuthorID:   author.UserID,
		Visibility: util.RandomInt32(0, 3),
		Data:       util.RandomRecipeData(),
	}

	recipe, err := testQueries.CreateRecipe(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, recipe)
	require.Equal(t, args.AuthorID, recipe.AuthorID)
	require.Equal(t, args.Visibility, recipe.Visibility)
	require.NotZero(t, recipe.CreatedAt)
	require.NotZero(t, recipe.UpdatedAt)

	var argsData, recipeData util.RecipeData
	err = json.Unmarshal(args.Data, &argsData)
	require.NoError(t, err)
	err = json.Unmarshal(recipe.Data, &recipeData)
	require.NoError(t, err)
	require.Equal(t, argsData, recipeData)

	return recipe
}

func CreateRandomRecipes(t *testing.T, num int) []Recipe {
	author := CreateRandomUser(t)

	var recipes []Recipe
	for i := 0; i < num; i++ {
		args := CreateRecipeParams{
			AuthorID:   author.UserID,
			Visibility: util.RandomInt32(0, 3),
			Data:       util.RandomRecipeData(),
		}
		recipe, err := testQueries.CreateRecipe(context.Background(), args)
		require.NoError(t, err)
		require.NotEmpty(t, recipe)
		require.Equal(t, args.AuthorID, recipe.AuthorID)
		require.Equal(t, args.Visibility, recipe.Visibility)
		require.NotZero(t, recipe.CreatedAt)
		require.NotZero(t, recipe.UpdatedAt)

		var argsData, recipeData util.RecipeData
		err = json.Unmarshal(args.Data, &argsData)
		require.NoError(t, err)
		err = json.Unmarshal(recipe.Data, &recipeData)
		require.NoError(t, err)
		require.Equal(t, argsData, recipeData)

		recipes = append(recipes, recipe)
	}
	return recipes
}

func TestCreateRecipe(t *testing.T) {
	CreateRandomRecipe(t)
}

func TestDeleteRecipe(t *testing.T) {
	recipe1 := CreateRandomRecipe(t)
	_, err := testQueries.DeleteRecipe(context.Background(), recipe1.ID)

	recipe2, err2 := testQueries.FetchRecipe(context.Background(), recipe1.ID)
	require.NoError(t, err)
	require.Error(t, err2)
	require.EqualError(t, err2, "no rows in result set")
	require.Empty(t, recipe2)
}

func TestFetchRecipe(t *testing.T) {
	recipe1 := CreateRandomRecipe(t)

	recipe2, err := testQueries.FetchRecipe(context.Background(), recipe1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, recipe2)
	require.Equal(t, recipe1.ID, recipe2.ID)
	require.Equal(t, recipe1.AuthorID, recipe2.AuthorID)
	require.Equal(t, recipe1.Visibility, recipe2.Visibility)
	require.Equal(t, recipe1.Data, recipe2.Data)
	require.NotZero(t, recipe2.CreatedAt)
	require.NotZero(t, recipe2.UpdatedAt)
}

func TestListRecipesByAuthor(t *testing.T) {
	recipes := CreateRandomRecipes(t, 10)

	author := recipes[0].AuthorID

	arg := ListRecipesByAuthorParams{
		AuthorID: author,
		Limit:    5,
		Offset:   5,
	}

	recipes2, err := testQueries.ListRecipesByAuthor(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, recipes2, 5)

	for _, recipe := range recipes2 {
		require.NotEmpty(t, recipe)
		require.Equal(t, author, recipe.AuthorID)
	}
}

func TestListRecipes(t *testing.T) {
	CreateRandomRecipes(t, 10)

	arg := ListRecipesParams{
		Limit:  5,
		Offset: 5,
	}

	recipes2, err := testQueries.ListRecipes(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, recipes2, 5)

	for _, recipe := range recipes2 {
		require.NotEmpty(t, recipe)
	}
}

func TestUpdateRecipe(t *testing.T) {
	recipe1 := CreateRandomRecipe(t)

	args := UpdateRecipeParams{
		Visibility: util.RandomInt32(0, 3),
		Data:       util.RandomRecipeData(),
		ID:         recipe1.ID,
	}

	recipe2, err := testQueries.UpdateRecipe(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, recipe2)
	// require.NotEqual(t, recipe2.Visibility, recipe1.Visibility)
	require.NotEqual(t, recipe2.Data, recipe1.Data)
	require.NotZero(t, recipe2.CreatedAt)
	require.NotZero(t, recipe2.UpdatedAt)
	require.Equal(t, recipe1.ID, recipe2.ID)

	var argsData, recipeData util.RecipeData
	err = json.Unmarshal(args.Data, &argsData)
	require.NoError(t, err)
	err = json.Unmarshal(recipe2.Data, &recipeData)
	require.NoError(t, err)
	require.Equal(t, argsData, recipeData)
}
