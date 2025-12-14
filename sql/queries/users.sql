-- name: CreateUser: one
INSERT INTO USERS (id, created_at, updated_at, name, email, password)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;