-- name: GetUsers :many
SELECT * from users;

-- name: GetUser :one
SELECT ID,username,name,email from users 
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

-- name: GetBooks :many 
SELECT * from books
LIMIT $1
OFFSET $2;

-- name: BuyBook :one
INSERT into orders(
    userId,bookId
)VALUES($1,$2)RETURNING *;

-- name: GetOrders :many

SELECT id,bookId from orders WHERE(userId = $1);