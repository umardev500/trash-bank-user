package service

import (
	"context"
	"user/pb/user"
)

func (u *userService) FindOne(ctx context.Context, req *user.UserFindOneRequest) (res *user.UserFindOneResponse, err error) {
	item, err := u.repo.FindOne(ctx, req.UserId, req.User, req.Pass, req.IsLogin)
	if err != nil {
		return
	}

	if item == nil {
		res.IsEmpty = true
		err = nil
		return
	}

	res = &user.UserFindOneResponse{
		Payload: u.parseUser(*item),
	}

	return
}
