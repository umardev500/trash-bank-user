package service

import (
	"context"
	"user/pb/user"
)

func (u *userService) Find(ctx context.Context, req *user.UserFindRequest) (res *user.UserFindResponse, err error) {
	res, err = u.repo.Find(ctx, req.Sort, req.Page, req.PerPage, req.Status, req.Search)
	return
}
