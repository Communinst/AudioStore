package handler

import (
	"AudioShare/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthorizationHandlerInterface interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

type DumpHandlerInterface interface {
	CreateDump(c *gin.Context)
}

type UserHandlerInterface interface {
	ObtainProfileById(c *gin.Engine)
	ObtainAllUsers(c *gin.Engine)
	RemoveUserById(c *gin.Engine)
}

type TrackHandlerInterface interface {
	UploadTrack(c *gin.Context)
	DownloadTrack(c *gin.Context)
}

type Handler struct {
	auth  AuthorizationHandlerInterface
	user  UserHandlerInterface
	dump  DumpHandlerInterface
	track TrackHandlerInterface
}

func NewHandler(srvc *service.Service) *Handler {
	return &Handler{
		auth: NewAuthHandler(srvc.Auth),
		dump: NewDumpHandler(srvc.Dump),
		//track:
	}
}
