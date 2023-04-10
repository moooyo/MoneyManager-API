package main

import (
	"MoneyManager/src/database"
	"MoneyManager/src/router"
	"MoneyManager/src/utility/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.UseRouters(app)
	database.Init()
	config := utils.GetConfig()
	_ = app.Listen(fmt.Sprintf(":%d", config.App.Port))
}
