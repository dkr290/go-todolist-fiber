package main

import (
	"github.com/dkr290/go-todolist-fiber/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	DBConnect *gorm.DB
)

func main() {

	app := fiber.New()
	handlers.InitDatabase()
	handlers.SetupRoutes(app)
	app.Listen("127.0.0.1:8080")
}
