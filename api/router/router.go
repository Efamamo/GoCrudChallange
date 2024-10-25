package router

import (
	"github.com/Efamamo/GoCrudChallange/api/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartRouter(pc controller.PersonController) {
	r := gin.Default()

	corsConfiguration := cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE", "PUT"},
		AllowHeaders:    []string{"Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length"},
	})

	r.Use(corsConfiguration)

	r.POST("/person", pc.Create)
	r.GET("/person", pc.GetAll)
	r.GET("/person/:id", pc.Get)
	r.PUT("/person/:id", pc.Update)
	r.DELETE("/person/:id", pc.Delete)

	r.Run()
}
