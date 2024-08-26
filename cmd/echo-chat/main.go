package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sadnamSakib/echo-chat/internal/app/router"
	"github.com/sadnamSakib/echo-chat/internal/config"
	"github.com/sadnamSakib/echo-chat/internal/db"
)

func main() {
	// Load application configuration
	config.LoadConfig()

	// Initialize MongoDB connection
	db.Connect()
	defer db.Disconnect() // Ensure MongoDB connection is closed when the application exits

	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} [${method}] ${uri} (${status})\n",
	}))
	e.Use(middleware.Recover())

	// Serve static files
	e.Static("/", "static")

	// Initialize routes
	router.InitRoutes(e)

	// Start the server
	log.Fatal(e.Start(":8080"))
}
