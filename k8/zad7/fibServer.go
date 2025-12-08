package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func fibb(i int) int {
	if i == 1 || i == 0 {
		return 1
	}
	return fibb(i-1) + fibb(i-2)
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "fibb_request_total",
	})
)

func main() {
	r := gin.Default()

	r.GET("/inc", func(context *gin.Context) {
		opsProcessed.Inc()
		context.Status(200)
	})

	r.GET("/", func(c *gin.Context) {
		_, err := http.Get("http://prometheus/inc")
		if err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
			return
		}
		opsProcessed.Inc()
		res := c.Request.URL.Query().Get("i")
		i, err := strconv.Atoi(res)
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"res": fibb(i),
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":80")
}
