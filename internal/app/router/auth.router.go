package router

import (
	"github.com/labstack/echo/v4"
	"github.com/sadnamSakib/echo-chat/internal/app/controller"
	"github.com/sadnamSakib/echo-chat/internal/app/middleware"
)

// InitAuthRoutes sets up the routes for authentication
func InitAuthRoutes(e *echo.Echo) {
	authGroup := e.Group("/auth")
	authGroup.POST("/login", controller.Login)
	authGroup.POST("/register", controller.Register)
	authGroup.Use(middleware.RedirectIfAuthenticated)
}
