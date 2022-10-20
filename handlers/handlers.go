package handlers

import (
	"fmt"

	"github.com/dkr290/go-todolist-fiber/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConnect *gorm.DB

var Repo = models.NewRepo(DBConnect)

func InitDatabase() {
	var err error
	dsn := "host=172.31.121.144 user=postgres password=Password123 dbname=goTodo port=5432"
	Repo.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	fmt.Println("Db connects")
	Repo.DB.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated")

}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/get_todos", Repo.GetTodos)
	api.Post("/create_todos", Repo.CreateTodo)
	api.Get("/get_todos/:id", Repo.GetTodoById)
	api.Put("/update_todos/:id", Repo.UpdateTodo)
	api.Delete("/delete_todos/:id", Repo.DeleteTodo)
}
