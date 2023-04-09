package login

import (
	"MoneyManager/src/database"
	"MoneyManager/src/utility/log"

	"github.com/gofiber/fiber/v2"
)

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(context *fiber.Ctx) {
	var param LoginParams
	if err := context.BodyParser(param); err != nil {
		log.TrackError(context, err)
		return
	}

	query := database.GetQuery()

	user, err := query.User.WithContext(context).FilterWithUsernameAndPassword(param.Username, param.Password)
	if err != nil {
		log.TrackError(context, err)
		context.JSON(400)
	}

	log.TrackTrace(context, "login success", "UserName", param.Username, "Password", param.Password)
}
