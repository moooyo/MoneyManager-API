package routers

import (
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	UseUserRouters(r.Group("/user"))
	UseLoginRouters(r.Group("/login"))
}
