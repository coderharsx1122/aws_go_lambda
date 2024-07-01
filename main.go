package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
)

func handler(ctx context.Context) (interface{}, error) {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from Lambda!"})
	})
	return r, nil
}

func main() {
	lambda.Start(handler)
}
