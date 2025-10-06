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
}

type TrackHandlerInteface interface {
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
