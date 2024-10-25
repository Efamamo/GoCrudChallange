package router

import (
	"github.com/Efamamo/GoCrudChallange/api/controller"
	"github.com/gin-gonic/gin"
)

func getHome(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"message": "Hello world"})
}

func StartRouter(pc controller.PersonController) {
	r := gin.Default()

	r.GET("/", getHome)
	r.POST("/person", pc.Create)

	r.Run()
}
