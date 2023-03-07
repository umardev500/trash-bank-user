package repo

import (
	"context"
	"user/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (u *userRepo) FindOne(ctx context.Context, userId, user, pass string, isLogin bool) (res *domain.User, err error) {
	filter := bson.M{"user_id": userId}
	if isLogin {
		filter = bson.M{"user": user, "pass": pass}
	}

	err = u.users.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return
	}

	return
}
