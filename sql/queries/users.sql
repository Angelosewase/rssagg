-- name: CreateUser :execresult
INSERT INTO users (
  id, createdAt, updateAt, name
)  VALUES(
  ?, ?, ?,?
);