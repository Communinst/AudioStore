package router

import (
	authToken "AudioShare/backend/internal/JSONWebTokens"
	"AudioShare/backend/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	dumps := v1.Group("/dumps")
	dumps.Use(authToken.JwtAuthMiddleware())
	{
		dumps.POST("/create", o.handler.Dump.CreateDump)
		dumps.POST("/restore", o.handler.Dump.RestoreDump)
		dumps.GET("/", o.handler.Dump.GetAllDumps)
	}
	return router
}
