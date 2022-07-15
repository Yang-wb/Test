package main

import (
	"math/rand"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	const keyRequestId = "requestId"

	//不管是get 还是post 都会优先进来 类似于拦截器
	r.Use(func(c *gin.Context) {
		s := time.Now()

		c.Next()
		// path, response code , log latency
		logger.Info("incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)))
	}, func(c *gin.Context) {
		c.Set(keyRequestId, rand.Int())
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {

		hs := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get(keyRequestId); exists {
			hs[keyRequestId] = rid
		}
		c.JSON(http.StatusOK, hs)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
