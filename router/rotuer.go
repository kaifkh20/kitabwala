package router

import (
	"kw/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/profile", handler.UserGet)
	user.Post("/create", handler.UserCreate)
	user.Post("/login", handler.UserLogin)

	book := app.Group("/book")
	book.Post("/addBook", handler.AddBook)
	book.Get("/books", handler.GetAllBook)
}
