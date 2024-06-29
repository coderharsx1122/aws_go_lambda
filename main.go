package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Serveless golang rest api",
		})
	})
	r.Run()
}
