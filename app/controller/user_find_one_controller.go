package controller

import (
	"context"
	"user/pb/user"
)

func (u *UserController) FindOne(ctx context.Context, req *user.UserFindOneRequest) (res *user.UserFindOneResponse, err error) {
	res, err = u.service.FindOne(ctx, req)
	return
}
