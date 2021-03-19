package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kazuki0924/go-fiber/book"
	"github.com/kazuki0924/go-fiber/database"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("hello world")
}

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/book/:id", book.GetBook)
	app.Get("api/v1/books", book.ListBooks)
	app.Post("api/v1/book", book.CreateBook)
	app.Patch("api/v1/book/:id", book.UpdateBook)
	app.Delete("api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}
