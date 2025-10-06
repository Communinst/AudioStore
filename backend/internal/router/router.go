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

	v1 := router.Group("/v1")
	auth := v1.Group("/auth")
	{
		auth.POST("/signup", o.handler.Auth.SignUp)
		auth.POST("/signin", o.handler.Auth.SignIn)
	}
	tracks := v1.Group("/tracks")
	{
		tracks.POST("/upload", o.handler.Track.UploadTrack)
		tracks.GET("/download/:id", o.handler.Track.DownloadTrack)
		tracks.GET("/info/:id", o.handler.Track.GetTrackInfo)
	}
	users := v1.Group("/users")
	{
		users.GET("/:id", o.handler.User.ObtainProfileById)
		users.GET("/", o.handler.User.ObtainAllUsers)
		users.DELETE("/:id", o.handler.User.RemoveUserById)
	}
	return router
}
