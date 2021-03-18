package book

import "github.com/gofiber/fiber"

func GetBook(c *fiber.Ctx) {
	c.Send("A Single Book")
}

func ListBooks(c *fiber.Ctx) {
	c.Send("All Books")
}

func CreateBook(c *fiber.Ctx) {
	c.Send("Creates a book")
}

func UpdateBook(c *fiber.Ctx) {
	c.Send("Updates a book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Deletes a book")
}
