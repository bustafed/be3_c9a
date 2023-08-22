package main

import "github.com/gin-gonic/gin"

func main() {
	// Crea un router con gin
	router := gin.Default()

	// Captura la solicitud GET “/hello-world”
	router.GET("/hola-mundo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensaje": "hola mundo",
		})
	})

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
