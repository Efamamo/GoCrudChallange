package router

import (
	"github.com/Efamamo/GoCrudChallange/api/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// StartRouter initializes the Gin router, sets up CORS, defines route handlers,
// and starts the HTTP server.
func StartRouter(pc controller.PersonController) {
	// Initialize a new Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Configure CORS settings to allow requests from any origin, specify allowed methods and headers.
	corsConfiguration := cors.New(cors.Config{
		AllowAllOrigins: true,                                     // Allow requests from all origins
		AllowMethods:    []string{"GET", "POST", "DELETE", "PUT"}, // Allowed HTTP methods
		AllowHeaders:    []string{"Origin", "Content-Type"},       // Allowed HTTP headers
		ExposeHeaders:   []string{"Content-Length"},               // Headers exposed to the client
	})

	// Apply the CORS middleware to all routes
	r.Use(corsConfiguration)

	// Define route handlers for Person entity CRUD operations
	r.POST("/person", pc.Create)       // Route for creating a new person
	r.GET("/person", pc.GetAll)        // Route for retrieving all persons
	r.GET("/person/:id", pc.Get)       // Route for retrieving a specific person by ID
	r.PUT("/person/:id", pc.Update)    // Route for updating an existing person by ID
	r.DELETE("/person/:id", pc.Delete) // Route for deleting a person by ID

	// Handler for undefined routes (404 Not Found)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error": "Route not found", // Response message for non-existing routes
		})
	})

	// Start the HTTP server on the default port (8080)
	r.Run()
}
