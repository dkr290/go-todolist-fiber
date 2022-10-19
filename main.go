package main

import (
	"fmt"

	"github.com/dkr290/go-todolist-fiber/database"
	"github.com/dkr290/go-todolist-fiber/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	app := fiber.New()
	initDatabase()
	app.Get("/", helloWorld)
	setupRoutes(app)
	app.Listen("127.0.0.1:8080")
}

func initDatabase() {
	var err error
	dsn := "host=192.168.105.133 user=postgres password=Password123 dbname=goTodo port=5432"
	database.DBConnect, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	fmt.Println("Db connectes")
	database.DBConnect.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated")

}

func setupRoutes(app *fiber.App) {

	app.Get("/todos", models.GetTodos)
	app.Post("/todos", models.CreateTodo)
	app.Get("/todos/:id", models.GetTodoById)
	app.Put("/todos/:id", models.UpdateTodo)
	app.Delete("/todos/:id", models.DeleteTodo)
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")

}
