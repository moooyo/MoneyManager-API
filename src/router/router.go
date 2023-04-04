package router

import (
	"github.com/gin-gonic/gin"
)

func UseRouters(r *gin.Engine) {
	useUserRouters(r.Group("/user"))
	useLoginRouters(r.Group("/login"))
}
