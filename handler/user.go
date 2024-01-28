package handler

import (
	"fmt"
	"kw/auth"
	"kw/database"
	"kw/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UserGet(c *fiber.Ctx) error {
	// token := c.Cookies("user_token")

	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	queries := model.New(database.DB)

	user, err := queries.GetUser(c.Context(), id)

	if err != nil {
		return err
	}

	return c.JSON(user)

	// return c.JSON(fiber.Map{
	// 	"token": token,
	// })
}

func UserLogin(c *fiber.Ctx) error {
	type UserLogin struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type Response struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	}

	payload := new(UserLogin)
	err := c.BodyParser(&payload)

	if err != nil {
		return err
	}

	// fmt.Println(payload.Email, payload.Password, payload.Username)

	queries := model.New(database.DB)

	user, err := queries.CheckLoginUser(c.Context(), payload.Username)

	fmt.Println(user)

	if err != nil {
		return err
	}

	passwordCheck := auth.CheckPasswordHash(payload.Password, user.Password)
	if !passwordCheck {
		return fiber.NewError(404, "Incorrect Password")
	}

	token, err := auth.GenerateToken(user.Username, user.Email)

	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:  "user_token",
		Value: token,
	})

	response := new(Response)
	response.Username = user.Username
	response.Email = user.Email
	response.Token = token

	return c.JSON(response)
}

func UserCreate(c *fiber.Ctx) error {
	type User struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	queries := model.New(database.DB)

	payload := new(User)
	err := c.BodyParser(&payload)

	// fmt.Println(payload.Email, payload.Password, payload.Username)

	if err != nil {
		return err
	}

	hashedPassword, err := auth.HashPassword(payload.Password)

	if err != nil {
		return c.Status(400).SendString("Error Occured")
	}

	insertedRow, err := queries.CreateUsers(c.Context(), model.CreateUsersParams{
		Name:     payload.Name,
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return err
	}

	// token:=auth.GenerateToken(payload.Username,)

	return c.JSON(insertedRow)

}
