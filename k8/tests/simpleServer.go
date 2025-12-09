package main

import "github.com/gin-gonic/gin"

type Res struct {
	Version int    `json:"version"`
	Id      string `json:"id"`
}

func main() {

	r := gin.Default()

	result := "hehe"

	r.GET("/test", func(c *gin.Context) {
		/*if rand.Int()%10 == 0 {
			os.Exit(1)
		}
		if rand.Int()%10 == 1 {
			for i := 1; i < 10; i++ {
				time.Sleep(time.Second)
			}
		}*/
		c.JSON(200, Res{
			Version: 2,
			Id:      result,
		})
	})

	r.GET("/setStr", func(context *gin.Context) {
		result = "not hehe"
	})

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run(":8080")
}
