package main

import (
	"MoneyManager/src/database"
	"MoneyManager/src/router"

	"github.com/gofiber/fiber/v2"
)

// use gin to create a web server
func main() {
	app := fiber.New()
	router.UseRouters(app)
	database.InitDB()
	app.Listen(":30000")
}
