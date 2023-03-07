package service

import (
	"context"
	"user/pb/user"

	"go.mongodb.org/mongo-driver/mongo"
)

func (u *userService) FindOne(ctx context.Context, req *user.UserFindOneRequest) (res *user.UserFindOneResponse, err error) {
	res = &user.UserFindOneResponse{}
	item, err := u.repo.FindOne(ctx, req.UserId, req.User, req.Pass, req.IsLogin)
	if err == mongo.ErrNoDocuments {
		res.IsEmpty = true
		err = nil
		return
	}

	if err != nil {
		return
	}

	if item == nil {
		res.IsEmpty = true
		err = nil
		return
	}

	res.Payload = u.parseUser(*item)

	return
}
