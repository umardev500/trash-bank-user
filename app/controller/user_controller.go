package controller

import (
	"user/domain"
	"user/pb/user"
)

type UserController struct {
	service domain.UserService
	user.UnimplementedUserServiceServer
}

func NewUserController(service domain.UserService) *UserController {
	return &UserController{service: service}
}
