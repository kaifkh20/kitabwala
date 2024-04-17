package handler

import (
	"kw/database"
	"kw/middleware"
	"kw/model"

	"github.com/gofiber/fiber/v2"
)

func AddBook(c* fiber.Ctx) error{
	_,err:= middleware.Protected(c)

	if(err!=nil){
		return err
	}

	queries :=
}