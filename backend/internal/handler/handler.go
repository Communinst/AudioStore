package handler

import "AudioShare/backend/internal/service"

type AuthorizationHandlerInterface interface {
	SignIn()
	SignUp()
}

type Handler struct {
	authorization AuthorizationHandlerInterface
}

func NewHandler(srvc *service.Service) *Handler {

}
