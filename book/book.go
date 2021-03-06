package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/kazuki0924/go-fiber/database"
)

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func ListBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func CreateBook(c *fiber.Ctx) {
	db := database.DBConn

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&book)
	c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	body := new(Book)

	if err := c.BodyParser(body); err != nil {
		c.Status(503).Send(err)
		return
	}

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book found with given ID")
		return
	}

	db.Model(&book).Updates(body)

	c.Send("Book successfully updated")
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book found with given ID")
		return
	}

	db.Delete(&book)
	c.Send("Book successfully deleted")
}
