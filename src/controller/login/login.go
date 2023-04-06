package login

import (
	"MoneyManager/src/utility/log"

	"github.com/gin-gonic/gin"
)

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(context *gin.Context) {
	var param LoginParams
	if err := context.Bind(param); err != nil {
		log.TrackError(context, err)
		return
	}

	log.TrackTrace(context, "login success", "UserName", param.Username, "Password", param.Password)
}
