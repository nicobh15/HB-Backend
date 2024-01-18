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
LIMIT $2
OFFSET $3;

-- name: ListRecipes :many
SELECT * FROM recipes
LIMIT $1
OFFSET $2;

-- name: UpdateRecipe :one
UPDATE recipes
SET 
    author_id = COALESCE($1, author_id),
    visibility = COALESCE($2, visibility),
    data = COALESCE($3, data)
WHERE id = $4
RETURNING *;

-- name: DeleteRecipe :one
DELETE FROM recipes
WHERE id = $1
RETURNING *;
