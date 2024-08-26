package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sadnamSakib/echo-chat/internal/config"
)

type jwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func VerifyJWTMiddleware() echo.MiddlewareFunc {
	secret := config.AppConfig.JWT.Secret
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey:  []byte(secret),
		TokenLookup: "cookie:token",
	}
	return echojwt.WithConfig(config)
}

func RedirectIfAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Retrieve the JWT token from the cookie
		cookie, err := c.Cookie("token")
		if err != nil || cookie == nil || cookie.Value == "" {
			// If there's no token or an error occurred, proceed with the request
			return next(c)
		}

		// Parse and validate the JWT token
		tokenString := cookie.Value
		secret := []byte(config.AppConfig.JWT.Secret)
		token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return secret, nil
		})

		// If the token is valid, redirect to the home page
		if err == nil && token.Valid {
			return c.Redirect(http.StatusFound, "/home")
		}

		// If token is invalid or any error occurred, proceed with the request
		return next(c)
	}
}
