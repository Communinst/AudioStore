package router

import (
	"AudioShare/backend/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title AudioShare API
// @version 1.0
// @description This is the AudioShare backend server API documentation
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /v1
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
		// @Summary User registration
		// @Description Register a new user account
		// @Tags auth
		// @Accept json
		// @Produce json
		// @Param request body entity.SignUpRequest true "User registration data"
		// @Success 201 {object} entity.SignUpResponse
		// @Failure 400 {object} error.Response
		// @Failure 409 {object} error.Response
		// @Failure 500 {object} error.Response
		// @Router /v1/auth/signup [post]
		auth.POST("/signup", o.handler.Auth.SignUp)

		// @Summary User login
		// @Description Authenticate user and return tokens
		// @Tags auth
		// @Accept json
		// @Produce json
		// @Param request body entity.SignInRequest true "User login credentials"
		// @Success 200 {object} entity.SignInResponse
		// @Failure 400 {object} error.Response
		// @Failure 401 {object} error.Response
		// @Failure 500 {object} error.Response
		// @Router /v1/auth/signin [post]
		auth.POST("/signin", o.handler.Auth.SignIn)
	}

	tracks := v1.Group("/tracks")
	{
		// @Summary Upload track
		// @Description Upload a new audio track
		// @Tags tracks
		// @Accept multipart/form-data
		// @Produce json
		// @Security ApiKeyAuth
		// @Param file formData file true "Audio file to upload"
		// @Param title formData string true "Track title"
		// @Param description formData string false "Track description"
		// @Success 201 {object} entity.Track
		// @Failure 400 {object} error.Response
		// @Failure 401 {object} error.Response
		// @Failure 413 {object} error.Response
		// @Failure 500 {object} error.Response
		// @Router /v1/tracks/upload [post]
		tracks.POST("/upload", o.handler.Track.UploadTrack)

		// @Summary Download track
		// @Description Download a track by ID
		// @Tags tracks
		// @Produce audio/*
		// @Param id path string true "Track ID"
		// @Success 200 {file} binary
		// @Failure 400 {object} error.Response
		// @Failure 404 {object} error.Response
		// @Failure 500 {object} error.Response
		// @Router /v1/tracks/download/{id} [get]
		tracks.GET("/download/:id", o.handler.Track.DownloadTrack)

		// @Summary Get track info
		// @Description Get track information by ID
		// @Tags tracks
		// @Produce json
		// @Param id path string true "Track ID"
		// @Success 200 {object} entity.Track
		// @Failure 400 {object} error.Response
		// @Failure 404 {object} error.Response
		// @Failure 500 {object} error.Response
		// @Router /v1/tracks/info/{id} [get]
		tracks.GET("/info/:id", o.handler.Track.GetTrackInfo)
	}

	users := v1.Group("/users")
	{
		// @Summary Get user by ID
		// @Description Get user profile information by user ID
		// @Tags users
		// @Produce json
		// @Security ApiKeyAuth
		// @Param id path string true "User ID"
		// @Success 200 {object} entity.User
		// @Failure 400 {object} error.Response
		// @Failure 401 {object} error.Response
		// @Failure 404 {object} error.Response
		// @Failure 500 {object} error.Response
		// @Router /v1/users/{id} [get]
		users.GET("/:id", o.handler.User.ObtainProfileById)

		// @Summary Get all users
		// @Description Get a list of all users (admin only)
		// @Tags users
		// @Produce json
		// @Security ApiKeyAuth
		// @Success 200 {array} entity.User
		// @Failure 401 {object} error.Response
		// @Failure 403 {object} error.Response
		// @Failure 500 {object} error.Response
		// @Router /v1/users/ [get]
		users.GET("/", o.handler.User.ObtainAllUsers)

		// @Summary Delete user by ID
		// @Description Delete a user by ID (admin only)
		// @Tags users
		// @Security ApiKeyAuth
		// @Param id path string true "User ID"
		// @Success 204
		// @Failure 400 {object} error.Response
		// @Failure 401 {object} error.Response
		// @Failure 403 {object} error.Response
		// @Failure 404 {object} error.Response
		// @Failure 500 {object} error.Response
		// @Router /v1/users/{id} [delete]
		users.DELETE("/:id", o.handler.User.RemoveUserById)
	}

	return router
}
