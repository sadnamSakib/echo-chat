package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware - Example middleware function for authentication
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Implement authentication logic here
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing authorization header"})
		}

		// If authenticated, call next handler
		return next(c)
	}
}
