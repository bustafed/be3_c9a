package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/hola-mundo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensaje": "hola mundo",
		})
	})

	router.Run()
}
