package repo

import (
	"user/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	users *mongo.Collection
}

func NewUserRepo(users *mongo.Collection) domain.UserRepo {
	return &userRepo{users: users}
}
