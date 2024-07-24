-- name: CreateUser :execresult
INSERT INTO users(
  id, created_at, updated_at, name, api_key
)  VALUES(
  ?, ?, ?,?,?
);

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key= ?;