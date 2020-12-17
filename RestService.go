package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	var profile BaseProfile

	router.POST("/register", func(c *gin.Context) {

		id := c.Query("id")
		dni := c.Query("dni")
		name := c.Query("name")
		gender := c.Query("gender")
		age := c.Query("email")
		address := c.Query("address")
		username := c.Query("username")
		password := c.Query("password")

	})


}



