// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package model

type Book struct {
	ID          int64
	Name        string
	Price       int32
	Description string
	Sellername  string
	Condition   bool
}

type User struct {
	ID       int64
	Name     string
	Username string
	Email    string
	Password string
}
