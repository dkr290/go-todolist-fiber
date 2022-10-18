package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New()
	app.Get("/", helloWorld)
	app.Listen("127.0.0.1:8080")
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")

}
