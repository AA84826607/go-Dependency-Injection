package router

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	// Creates a router without any middleware by default
	r := gin.New()
	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//apiV1 := r.Group("/api/v1")
	//users


	gin.SetMode("debug")
	return r
}
