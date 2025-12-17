package main

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()

	id := uuid.New()

	r.GET("/", func(c *gin.Context) {

		res := rand.Int()
		statusCode := 200
		if res%5 == 0 {
			statusCode = 500
		}
		c.JSON(statusCode, gin.H{
			"id": id,
		})
	})

	r.Run(":80")
}
