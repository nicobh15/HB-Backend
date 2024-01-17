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
SET 
    username = COALESCE(NULLIF($1, ''), username),
    email = COALESCE(NULLIF($2, ''), email),
    first_name = COALESCE(NULLIF($3, ''), first_name),
    password_hash = COALESCE(NULLIF($4, ''), password_hash),
    role = COALESCE(NULLIF($5, ''), role),
    household_id = COALESCE($6, household_id),
    updated_at = now()
WHERE user_id = $7
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users 
WHERE email = $1 
RETURNING *;