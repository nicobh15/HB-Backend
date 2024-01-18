-- name: CreateInventoryItem :one
INSERT INTO inventory (
    household_id, category, name, quantity, location
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: FetchInventoryItem :one
SELECT * FROM inventory 
WHERE item_id = $1
LIMIT 1;

-- name: ListInventoryItems :many
SELECT * FROM inventory
WHERE household_id = $1
LIMIT $2
OFFSET $3;

-- name: UpdateInventoryItem :one
UPDATE inventory
SET 
    household_id = COALESCE($1, household_id), 
    category = COALESCE($2, category), 
    name = COALESCE($3, name), 
    quantity = COALESCE($4, quantity), 
    expiration_date = COALESCE($5, expiration_date), 
    purchase_date = COALESCE($6, purchase_date), 
    location = COALESCE($7, location)
WHERE item_id = $8
RETURNING *;

-- name: DeleteInventoryItem :one
DELETE FROM inventory
WHERE item_id = $1
RETURNING *;

-- name: ListInventoryItemsByCategory :many
SELECT * FROM inventory
WHERE household_id = $1 AND category = $2
LIMIT $3
OFFSET $4;

-- name: ListInventoryItemsByLocation :many
SELECT * FROM inventory
WHERE household_id = $1 AND location = $2
LIMIT $3
OFFSET $4;
