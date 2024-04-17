-- name: GetUsers :many
SELECT * from users;

-- name: GetUser :one
SELECT * from users 
WHERE email = $1;

-- name: CheckLoginUser :one
SELECT username,email,password from users
WHERE username = $1 ;


-- name: CreateUsers :one 
INSERT into users (
    name,username,email,password
)   
VALUES(
        $1, $2, $3, $4
    )
RETURNING *;


-- name: AddBooks :one 
INSERT into books(
    name,price,description,sellerName,condition
)
VALUES(
    $1,$2,$3,$4,$5
)RETURNING *;