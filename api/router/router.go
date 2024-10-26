package router

import (
	"fmt"

	"github.com/Efamamo/GoCrudChallange/api/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	host        string
	port        string
	controllers []any
}

// Config holds configuration settings for creating a new Router instance.
type Config struct {
	Host        string
	Port        string
	Controllers []any // List of controllers
}

// NewRouter creates a new Router instance with the given configuration.
func NewRouter(config Config) *Router {
	return &Router{
		port:        config.Port,
		host:        config.Host,
		controllers: config.Controllers,
	}
}

// StartRouter initializes the Gin router, sets up CORS, defines route handlers, and starts the HTTP server.
func (router *Router) StartRouter(pc controller.PersonController) {
	r := gin.Default()

	// Configure CORS settings to allow requests from any origin, specify allowed methods and headers.
	corsConfiguration := cors.New(cors.Config{
		AllowAllOrigins: true,                                     // Allow requests from all origins
		AllowMethods:    []string{"GET", "POST", "DELETE", "PUT"}, // Allowed HTTP methods
		AllowHeaders:    []string{"Origin", "Content-Type"},       // Allowed HTTP headers
		ExposeHeaders:   []string{"Content-Length"},               // Headers exposed to the client
	})

	r.Use(corsConfiguration)

	// Group all routes related to person operations
	personRoutes := r.Group("/person")
	{
		personRoutes.POST("", pc.Create)       // POST /person
		personRoutes.GET("", pc.GetAll)        // GET /person
		personRoutes.GET("/:id", pc.Get)       // GET /person/:id
		personRoutes.PUT("/:id", pc.Update)    // PUT /person/:id
		personRoutes.DELETE("/:id", pc.Delete) // DELETE /person/:id
	}

	// Handler for undefined routes (404 Not Found)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error": "Route not found",
		})
	})

	r.Run(fmt.Sprintf("%s:%s", router.host, router.port))
}
