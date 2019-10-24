package main

import (
	"fmt"
)

import (
	"github.com/gin-gonic/gin"
	"github.com/fvbock/endless"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
				"message": "pong",
		})
	})

	router.POST("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
				"message": "pong",
		})
	})

	return router
}


func main() {
	fmt.Println("Test Server")
	endless.ListenAndServe(":8080", setUpRouter())
}

