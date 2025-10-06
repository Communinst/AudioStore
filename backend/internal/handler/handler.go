package handler

import (
	"AudioShare/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthorizationHandlerInterface interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

type UserHandlerInterface interface {
	ObtainProfileById(c *gin.Context)
	ObtainAllUsers(c *gin.Context)
	RemoveUserById(c *gin.Context)
}

type TrackHandlerInterface interface {
	UploadTrack(c *gin.Context)
	DownloadTrack(c *gin.Context)
	GetTrackInfo(c *gin.Context)
}

type Handler struct {
	Auth  AuthorizationHandlerInterface
	User  UserHandlerInterface
	Track TrackHandlerInterface
}

func NewHandler(srvc *service.Service) *Handler {
	return &Handler{
		Auth:  NewAuthHandler(srvc.AuthServiceInterface),
		User:  NewUserHandler(srvc.UserServiceInterface),
		Track: NewTrackHandler(srvc.TrackServiceInterface),
	}
}
