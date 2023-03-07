package service

import (
	"context"
	"math"
	"user/pb/user"
)

func (u *userService) Find(ctx context.Context, req *user.UserFindRequest) (res *user.UserFindResponse, err error) {
	sort := req.Sort
	page := req.Page
	perPage := req.PerPage
	if sort == 0 {
		sort = 1
	}
	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}

	// do finding
	items, rows, err := u.repo.Find(ctx, req.Sort, req.Page, req.PerPage, req.Status, req.Search)
	res = &user.UserFindResponse{} // init response

	// check for items
	if items == nil {
		res.IsEmpty = true
		return
	}

	var users []*user.User
	for _, val := range items {
		item := u.parseUser(val)
		users = append(users, item)
	}

	payload := &user.UserFind{
		Users: users,
	}
	pages := 1
	if rows >= perPage {
		pages = int(math.Ceil(float64(rows) / float64(perPage)))
	}
	payload.Rows = rows
	payload.Pages = int64(pages)
	payload.PerPage = perPage
	payload.Total = int64(len(users))
	res.Payload = payload

	return
}
