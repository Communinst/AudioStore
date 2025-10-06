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
	ObtainProfileById(c *gin.Engine)
	ObtainAllUsers(c *gin.Engine)
	RemoveUserById(c *gin.Engine)
}

type TrackHandlerInteface interface {
	UploadTrack(c *gin.Context)
	DownloadTrack(c *gin.Context)
}

type Handler struct {
	auth  AuthorizationHandlerInterface
	user  UserHandlerInterface
	track TrackHandlerInteface
}

func NewHandler(srvc *service.Service) *Handler {
	return &Handler{
		auth: NewAuthHandler(srvc.AuthServiceInteface),
	}
}
