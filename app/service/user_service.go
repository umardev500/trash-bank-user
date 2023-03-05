package service

import "user/domain"

type userService struct {
	repo domain.UserRepo
}

func NewUserService(repo domain.UserRepo) domain.UserService {
	return &userService{repo: repo}
}
