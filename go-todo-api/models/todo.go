package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB     *gorm.DB
	DbHost string
	DbUser string
	DbPass string
	DbPort string
	DbName string
}

// creates new repository
func NewRepo(d *gorm.DB) *Repository {
	r := Repository{
		DB: d,
	}
	return &r
}

type Todo struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (r *Repository) GetTodos(c *fiber.Ctx) error {
	db := r.DB
	var todos []Todo
	db.Find(&todos)
	return c.JSON(&todos)
}

func (r *Repository) CreateTodo(c *fiber.Ctx) error {
	db := r.DB
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	db.Create(&todo)

	return c.Status(200).JSON(todo)

}

func (r *Repository) GetTodoById(c *fiber.Ctx) error {

	db := r.DB
	var todos []Todo
	id := c.Params("id")
	if err := db.Find(&todos, id).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(&todos)

}

func (r *Repository) UpdateTodo(c *fiber.Ctx) error {
	type UpdatedTodo struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	db := r.DB
	var todos Todo
	id := c.Params("id")
	if err := db.Find(&todos, id).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var updTodo UpdatedTodo
	if err := c.BodyParser(&updTodo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	todos.Title = updTodo.Title
	todos.Completed = updTodo.Completed
	db.Save(&todos)

	return c.Status(400).JSON(&todos)

}

func (r *Repository) DeleteTodo(c *fiber.Ctx) error {
	db := r.DB
	var todos Todo
	id := c.Params("id")
	if err := db.Find(&todos, id).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	db.Delete(&todos)

	return c.SendStatus(200)
}
