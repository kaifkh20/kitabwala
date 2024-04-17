package handler

import (
	"kw/database"
	"kw/middleware"
	"kw/model"

	"github.com/gofiber/fiber/v2"
)

func AddBook(c *fiber.Ctx) error {
	email, err := middleware.Protected(c)

	if err != nil {
		return err
	}

	type Book struct {
		Name        string `json:"name"`
		Price       int32  `json:"price"`
		Description string `json:"description"`
		SellerName  string `json:"sellerName"`
		Condition   bool   `json:"condition"`
	}

	queries := model.New(database.DB)

	payload := new(Book)
	err = c.BodyParser(&payload)

	if err != nil {
		return err
	}

	payload.Condition = false
	payload.SellerName = email

	book, err := queries.AddBooks(c.Context(), model.AddBooksParams{
		Name:        payload.Name,
		Price:       payload.Price,
		Description: payload.Description,
		Sellername:  payload.SellerName,
		Condition:   payload.Condition,
	})

	if err != nil {
		return err
	}

	return c.JSON(book)
}

func GetAllBook(c *fiber.Ctx) error {
	queries := model.New(database.DB)

	books, err := queries.GetBooks(c.Context())

	if err != nil {
		return err
	}

	return c.JSON(books)

}
