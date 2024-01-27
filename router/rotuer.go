package router

import (
	"kw/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/:id", handler.UserGet)
	user.Post("/create", handler.UserCreate)
	user.Post("/login", handler.UserLogin)
}
