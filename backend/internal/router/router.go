package router

import (
	"AudioShare/backend/internal/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	handler *handler.Handler
}

func NewRouter(h *handler.Handler) *Router {
	return &Router{
		handler: h,
	}
}

func (o *Router) InitNewRouter(middleware ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middleware...)

	// v1 := router.Group("/v1")
	// auth := v1.Group("/auth")
	// {
	// auth.POST("/sign-in", o.handler.authorization.SignIn())
	// auth.POST("/sing-up", o.handler.authorization.SignUp())
	// }

	return router
}
