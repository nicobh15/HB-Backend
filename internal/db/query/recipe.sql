-- name: CreateRecipe :one
INSERT INTO recipes (
    author_id, visibility, data
    ) VALUES ( 
        $1, $2, $3
    ) RETURNING *;

-- name: FetchRecipe :one
SELECT * FROM recipes 
WHERE id = $1
LIMIT 1;

-- name: ListRecipesByAuthor :many
SELECT * FROM recipes
WHERE author_id = $1
LIMIT $2;

-- name: ListRecipes :many
SELECT * FROM recipes
LIMIT $1;

-- name: UpdateRecipe :one
UPDATE recipes
SET author_id = $1, visibility = $2, data = $3
WHERE id = $4
RETURNING *;

-- name: DeleteRecipe :one
DELETE FROM recipes
WHERE id = $1
RETURNING *;
