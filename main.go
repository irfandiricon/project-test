package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware bawaan Echo
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Ambil port dari environment
	port := os.Getenv("WEB_PORT")
	if port == "" {
		port = "3000"
	}

	// Routing sederhana
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "Hello from Echo!",
			"status":  "running",
		})
	})

	log.Printf("🚀 Echo service running on port %s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
