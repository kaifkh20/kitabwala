package handler

import (
	"kw/database"
	"kw/middleware"
	"kw/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
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

	limitString := c.Params("limit", "10")

	limit, err := strconv.ParseInt(limitString, 10, 32)

	if err != nil {
		return err
	}

	offset, err := strconv.ParseInt(c.Params("offset", "0"), 10, 32)

	if err != nil {
		return err
	}

	books, err := queries.GetBooks(c.Context(), model.GetBooksParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})

	if err != nil {
		return err
	}

	return c.JSON(books)

}

func BuyBook(c *fiber.Ctx) error {

	email, err := middleware.Protected(c)

	if err != nil {
		return err
	}

	queries := model.New(database.DB)

	user, err := queries.GetUser(c.Context(), email)

	if err != nil {
		return err
	}

	type PayloadOrder struct {
		UserId pgtype.Int8 `json:"userId"`
		BookId pgtype.Int8 `json:"orderId"`
	}

	payload := new(PayloadOrder)
	err = c.BodyParser(&payload)

	if err != nil {
		return err
	}

	payload.UserId = pgtype.Int8{Int64: user.ID, Valid: true}

	book, err := queries.BuyBook(c.Context(), model.BuyBookParams{
		Userid: payload.UserId,
		Bookid: payload.BookId,
	})

	if err != nil {
		return err
	}

	return c.JSON(book)

}
