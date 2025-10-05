package handler

import (
	"AudioShare/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthorizationHandlerInterface interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

type Handler struct {
	authorization AuthorizationHandlerInterface
}

func NewHandler(srvc *service.Service) *Handler {
	return nil
}
