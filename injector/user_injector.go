package injector

import (
	"user/app/controller"
	"user/app/repo"
	"user/app/service"
	"user/conf"
)

func NewUserInjector() *controller.UserController {
	usersCollection := conf.NewConn().Collection("users")
	userRepo := repo.NewUserRepo(usersCollection)
	userService := service.NewUserService(userRepo)
	return controller.NewUserController(userService)
}
