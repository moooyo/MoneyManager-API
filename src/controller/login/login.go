package login

import (
	"MoneyManager/src/database"
	"MoneyManager/src/utility/log"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string
}

func Login(context *fiber.Ctx) {
	var param LoginParams
	if err := context.BodyParser(param); err != nil {
		log.TrackError(context.Context(), err)
		return
	}

	query := database.GetQuery()

	user, err := query.User.WithContext(context.Context()).FilterWithUsernameAndPassword(param.Username, param.Password)
	if err != nil {
		log.TrackError(context.Context(), err)
		if err := context.JSON(LoginResponse{
			Message: "Invalid username or password",
		}); err != nil {
			log.TrackError(context.Context(), err)
		}

		_ = context.SendStatus(400)
	}

	if len(user) > 0 {
		log.TrackTrace(context.Context(), "login success", "UID", fmt.Sprintf("%d", user[0].ID))
	}

	log.TrackTrace(context.Context(), "login success", "UserName", param.Username, "Password", param.Password)
}
