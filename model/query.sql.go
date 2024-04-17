// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package model

import (
	"context"
)

const addBooks = `-- name: AddBooks :one
INSERT into books(
    name,price,description,sellerName,condition
)
VALUES(
    $1,$2,$3,$4,$5
)RETURNING id, name, price, description, sellername, condition
`

type AddBooksParams struct {
	Name        string
	Price       int32
	Description string
	Sellername  string
	Condition   bool
}

func (q *Queries) AddBooks(ctx context.Context, arg AddBooksParams) (Book, error) {
	row := q.db.QueryRow(ctx, addBooks,
		arg.Name,
		arg.Price,
		arg.Description,
		arg.Sellername,
		arg.Condition,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Description,
		&i.Sellername,
		&i.Condition,
	)
	return i, err
}

const checkLoginUser = `-- name: CheckLoginUser :one
SELECT username,email,password from users
WHERE username = $1
`

type CheckLoginUserRow struct {
	Username string
	Email    string
	Password string
}

func (q *Queries) CheckLoginUser(ctx context.Context, username string) (CheckLoginUserRow, error) {
	row := q.db.QueryRow(ctx, checkLoginUser, username)
	var i CheckLoginUserRow
	err := row.Scan(&i.Username, &i.Email, &i.Password)
	return i, err
}

const createUsers = `-- name: CreateUsers :one
INSERT into users (
    name,username,email,password
)   
VALUES(
        $1, $2, $3, $4
    )
RETURNING id, name, username, email, password
`

type CreateUsersParams struct {
	Name     string
	Username string
	Email    string
	Password string
}

func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error) {
	row := q.db.QueryRow(ctx, createUsers,
		arg.Name,
		arg.Username,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name, username, email, password from users 
WHERE email = $1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, name, username, email, password from users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Username,
			&i.Email,
			&i.Password,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
