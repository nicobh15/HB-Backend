-- name: CreateHousehold :one
INSERT INTO households (
    household_name, address
    ) VALUES ( 
        $1, $2
    ) RETURNING *;

-- name: FetchHousehold :one
SELECT * FROM households 
WHERE household_id = $1
LIMIT 1;

-- name: UpdateHousehold :one
UPDATE households
SET household_name = $1, address = $2
WHERE household_id = $3
RETURNING *;

-- name: DeleteHousehold :one
DELETE FROM households
WHERE household_id = $1
RETURNING *;

-- name: ListHouseholds :many
SELECT * FROM households
LIMIT $1
OFFSET $2;

