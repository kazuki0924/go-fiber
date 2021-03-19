package book

import (
	"time"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/kazuki0924/go-fiber/database"
)

type Book struct {
	gorm.Model
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
}

func GetBook(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find((&books))
	c.JSON(books)
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
