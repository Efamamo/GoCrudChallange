package http

import "github.com/gin-gonic/gin"

func getHome(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"message": "Hello world"})
}

func StartRouter() {
	r := gin.Default()

	r.GET("/", getHome)

	r.Run()
}
