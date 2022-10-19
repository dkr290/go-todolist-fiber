package models

import (
	"github.com/dkr290/go-todolist-fiber/database"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetTodos(c *fiber.Ctx) error {
	db := database.DBConnect
	var todos []Todo
	db.Find(&todos)
	return c.JSON(&todos)
}

func CreateTodo(c *fiber.Ctx) error {
	db := database.DBConnect
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	db.Create(&todo)

	return c.Status(200).JSON(todo)

}

func GetTodoById(c *fiber.Ctx) error {

	db := database.DBConnect
	var todos []Todo
	id := c.Params("id")
	if err := db.Find(&todos, id).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(&todos)

}

func UpdateTodo(c *fiber.Ctx) error {
	type UpdatedTodo struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	db := database.DBConnect
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
