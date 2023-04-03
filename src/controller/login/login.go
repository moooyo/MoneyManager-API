package login

import "github.com/gin-gonic/gin"

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(context *gin.Context) {
	var param LoginParams
	if err := context.Bind(param) {
		context.JSON(200, gin.H{
			"message": "error",
		})
		return
	}
}
