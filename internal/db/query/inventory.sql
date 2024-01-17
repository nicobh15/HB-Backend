-- name: CreateInventoryItem :one
INSERT INTO inventory (
    household_id, category, name, quantity, expiration_date, purchase_date, location
    ) VALUES ( 
        $1, $2, $3, $4, $5, $6, $7
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
SET household_id = $1, category = $2, name = $3, quantity = $4, expiration_date = $5, purchase_date = $6, location = $7
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
