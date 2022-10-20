package main

import (
	"github.com/dkr290/go-todolist-fiber/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())
	handlers.InitDatabase()
	handlers.SetupRoutes(app)
	app.Listen("127.0.0.1:8080")
}
