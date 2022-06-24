package main

import (
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const keyrequestId = "requestId"

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	r.Use(func(c *gin.Context) {
		//Path,responsecode,log latency
		s := time.Now()
		c.Next()
		logger.Info("incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elaped", time.Now().Sub(s)))

	}, func(c *gin.Context) {
		//为request添加编号
		c.Set(keyrequestId, rand.Int())

		c.Next()
	})
	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get(keyrequestId); exists {
			h[keyrequestId] = rid
		}
		c.JSON(http.StatusOK, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
