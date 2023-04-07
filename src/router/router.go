package router

import (
	"github.com/gofiber/fiber/v2"
)

func UseRouters(app *fiber.App) {
	useUserRouters(app.Group("/user"))
	useLoginRouters(app.Group("/login"))
}
