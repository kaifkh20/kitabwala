package middleware

import (
	"fmt"
	"kw/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Protected(c *fiber.Ctx) (string, error) {
	// tokenStr := c.Cookies("user_token")

	tokenStr := c.Get("Authorization")
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ") // Access the Authorization header value
	if tokenStr == "" {
		return "", fiber.NewError(404, "Unauthorized Request")
	}

	// fmt.Println(tokenStr)

	// if tokenStr == "" {
	// 	return "", fiber.NewError(404, "Unauthorized Request")
	// }

	token, err := auth.AuthorizeToken(tokenStr)

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return fmt.Sprint(claims["email"]), nil
	}

	return "", fiber.NewError(500, "Server Error")
}
