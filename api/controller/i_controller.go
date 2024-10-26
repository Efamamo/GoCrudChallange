package controller

import "github.com/gin-gonic/gin"

// - Public: No authentication required.
// - Protected: Requires authentication.
// - Privileged: Requires additional role-based authorization.
type IController interface {
	// RegisterPublic sets up public routes.
	RegisterPublic(route *gin.RouterGroup)

	// RegisterProtected sets up routes that require authentication.
	RegisterProtected(route *gin.RouterGroup)

	// RegisterPrivileged sets up routes with additional role-based authorization.
	RegisterPrivileged(route *gin.RouterGroup)
}
