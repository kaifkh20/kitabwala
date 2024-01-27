-- name: GetUsers :many
SELECT * from users;

-- name: GetUser :one
SELECT * from users 
WHERE id = $1;

-- name: CheckLoginUser :one
SELECT username,email,password from users
WHERE username = $1 and email = $2;


-- name: CreateUsers :one 
INSERT into users (
    name,username,email,password
)   
VALUES(
        $1, $2, $3, $4
    )
RETURNING *;