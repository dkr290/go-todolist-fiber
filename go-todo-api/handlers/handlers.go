package handlers

import (
	"fmt"
	"os"

	"github.com/dkr290/go-todolist-fiber/go-todo-api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConnect *gorm.DB

var Repo = models.NewRepo(DBConnect)

func InitDatabase() {

	Repo.DbHost = os.Getenv("DATABASE_HOST")
	Repo.DbPass = os.Getenv("DATABASE_PASS")
	Repo.DbUser = os.Getenv("DATABASE_USER")
	Repo.DbPort = os.Getenv("DATABASE_PORT")
	Repo.DbName = os.Getenv("DB_NAME")
	var err error

	dsn := "host=" + Repo.DbHost + "user=" + Repo.DbUser + "password=" + Repo.DbPass + "dbname=" + Repo.DbName + "port=" + Repo.DbPort
	Repo.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	fmt.Println("Db connects")
	Repo.DB.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated")

}

// Another with local database
// func InitDatabase() {
// 	var err error

// 	Repo.DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to connect to the database")
// 	}

// 	fmt.Println("Db connects")
// 	Repo.DB.AutoMigrate(&models.Todo{})
// 	fmt.Println("Migrated")

// }

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/get_todos", Repo.GetTodos)
	api.Post("/create_todos", Repo.CreateTodo)
	api.Get("/get_todos/:id", Repo.GetTodoById)
	api.Put("/update_todos/:id", Repo.UpdateTodo)
	api.Delete("/delete_todos/:id", Repo.DeleteTodo)
}
