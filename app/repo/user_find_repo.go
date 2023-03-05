package repo

import (
	"context"
	"math"
	"user/domain"
	"user/helper"
	"user/pb/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (u *userRepo) Find(ctx context.Context, sort, page, perPage int64, status, search string) (res *user.UserFindResponse, err error) {
	if sort == 0 {
		sort = 1
	}
	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}

	res = &user.UserFindResponse{}

	filterData1 := helper.GetSerchRegex([]string{
		"user_id",
		"user",
		"details.name",
		"details.email",
		"details.phone.number",
		"details.address.province",
		"details.address.city",
		"details.address.district",
		"details.address.village",
		"details.address.postal_code",
	}, search)

	// filterData2 := []bson.M{}
	filter := bson.M{
		"$or": filterData1,
	}

	findOption := options.Find()

	// pagination
	offset := (page - 1) * perPage
	findOption.SetSkip(offset)
	findOption.SetLimit(perPage)
	// set sorting
	sortField := bson.M{"user_id": sort}
	findOption.SetSort(sortField)

	cur, err := u.users.Find(ctx, filter)
	if err != nil {
		return
	}

	var users []*user.User
	for cur.Next(ctx) {
		each := domain.User{}
		err := cur.Decode(&each)
		if err != nil {
			return nil, err
		}

		item := u.parseUser(each)
		users = append(users, item)
	}

	if len(users) < 1 {
		res.IsEmpty = true
		return
	}

	payload := &user.UserFind{
		Users: users,
	}

	// get all rows
	rows, _ := u.users.CountDocuments(ctx, filter)
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
