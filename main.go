package main

import (
	"github.com/gofiber/fiber"
	"github.com/kazuki0924/go-fiber/book"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("hello world")
}

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/book/:id", book.GetBook)
	app.Get("api/v1/books", book.ListBooks)
	app.Post("api/v1/book", book.CreateBook)
	app.Patch("api/v1/book", book.UpdateBook)
	app.Delete("api/v1/book", book.DeleteBook)
}

func main() {
	app := fiber.New()

	// app.Get("/", helloWorld)
	setupRoutes(app)

	app.Listen(3000)
}
