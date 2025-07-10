package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type content struct {
	Msg string `json:"msg"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})

	r.POST("/hook", func(c *gin.Context) {
		var content content

		if err := c.ShouldBindJSON(&content); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, content)
	})

	r.Run(":3000")
}
