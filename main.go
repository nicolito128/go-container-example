package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const PORT = ":8080"

func handler(c echo.Context) error {
	text := c.Param("*") // Catch-all parameter
	if text == "" {
		text = "Hello, World!"
	}

	return c.String(http.StatusOK, fmt.Sprintf("You requested: %s", text))
}

func main() {
	e := echo.New()

	// Add middleware to log all requests
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[ REQUEST ] method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/*", handler)

	fmt.Println("Starting server on http://localhost:" + PORT + "/ - Press CTRL+C to exit")
	if err := e.Start(PORT); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
