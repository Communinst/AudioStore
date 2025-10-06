package handler

import "AudioShare/backend/internal/service"

type UserHandler struct {
	srvc service.UserServiceInterface
}

func NewUserHandler(srvc service.UserServiceInterface) *UserHandler {
	return &UserHandler{
		srvc: srvc,
	}
}
