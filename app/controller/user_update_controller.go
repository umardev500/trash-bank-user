package controller

import (
	"context"
	"user/pb"
	"user/pb/user"
)

func (u *UserController) Update(ctx context.Context, req *user.UserUpdateRequest) (res *pb.OperationResponse, err error) {
	res, err = u.service.Update(ctx, req)
	return
}
