package controller

import (
	"net/http"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sadnamSakib/echo-chat/internal/config"
	"github.com/sadnamSakib/echo-chat/internal/repository"
	"github.com/sadnamSakib/echo-chat/pkg/models"
)

type jwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Login handles user login
func Login(c echo.Context) error {

	email := c.FormValue("email")
	password := c.FormValue("password")
	user, err := repository.FindUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
	}
	if err := repository.ComparePasswords(user.Password, password); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
	}
	claims := &jwtCustomClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := config.AppConfig.JWT.Secret
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = time.Now().Add(72 * time.Hour) // Set cookie expiration
	cookie.HttpOnly = true                          // Make the cookie HTTP-only
	cookie.Secure = false                           // Use Secure flag (set to true for HTTPS)
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{"message": "Logged in successfully"})

}

func Register(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: password,
		UserId:   uuid.New().String(),
	}
	err := repository.SaveUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to save user"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "User saved"})

}
