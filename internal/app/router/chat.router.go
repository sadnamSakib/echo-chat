package router

import (
	"github.com/labstack/echo/v4"
	"github.com/sadnamSakib/echo-chat/internal/app/controller"
	"github.com/sadnamSakib/echo-chat/internal/app/middleware"
)

// InitAuthRoutes sets up the routes for authentication
func InitChatRoutes(e *echo.Echo) {
	chatGroup := e.Group("/chat")
	chatGroup.GET("/messages/:roomId", controller.ReceiveMessages)
	chatGroup.POST("/messages/:roomId", controller.SendMessage)
	chatGroup.POST("/newchat", controller.NewChatRoom)
	chatGroup.GET("/add/:roomId", controller.AddUserToRoom)
	chatGroup.Use(middleware.VerifyJWTMiddleware())
}
