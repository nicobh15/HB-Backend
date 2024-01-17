-- name: CreateUser :one
INSERT INTO users (
    username, email,first_name, password_hash, role, household_id
    ) VALUES ( 
        $1, $2, $3, $4, $5, $6
    ) RETURNING *;

-- name: FetchUserByUserName :one
SELECT * FROM users 
WHERE username = $1
LIMIT 1;

-- name: FetchUserByUserId :one
SELECT * FROM users 
WHERE user_id = $1
LIMIT 1;

-- name: FetchUserByEmail :one
SELECT * FROM users 
WHERE email = $1
LIMIT 1;

-- name: ListHouseholdMembers :many
SELECT * FROM users 
WHERE household_id = $1
LIMIT $2
OFFSET $3;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username ASC
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users 
SET username = $1, email = $2, password_hash = $3, role = $4, household_id = $5, updated_at = now()
WHERE user_id = $6
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users 
WHERE email = $1 
RETURNING *;