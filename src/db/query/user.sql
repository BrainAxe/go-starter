-- name: CreateUser :one
INSERT INTO users (
  username,
  email,
  password,
  mobile
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
set email = $2,
mobile = $3
WHERE id = $1
RETURNING *;

-- name: ListUsers :many
SELECT username, email, mobile, created_at FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;