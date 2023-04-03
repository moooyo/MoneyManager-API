package main

import (
	"MoneyManager/src/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

// use gin to create a web server
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.UseRouters(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
