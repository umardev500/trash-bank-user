package controller

import (
	"context"
	"user/pb/user"
)

func (u *UserController) Find(ctx context.Context, req *user.UserFindRequest) (res *user.UserFindResponse, err error) {
	res, err = u.service.Find(ctx, req)

	return
}
