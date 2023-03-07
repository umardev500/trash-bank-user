package repo

import (
	"context"
	"user/domain"
	"user/helper"
	"user/variable"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (u *userRepo) Find(ctx context.Context, sort, page, perPage int64, status, search string) (res []domain.User, rows int64, err error) {
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

	deletedValue := bson.M{"$eq": nil} // default deleted status is not deleted
	if status == variable.StatusDeleted {
		deletedValue = bson.M{"$ne": nil}
	}
	deletedFilter := bson.M{"status.0.status_text": deletedValue} // deleted filter

	filterData2 := []bson.M{}
	filterData2 = append(filterData2, deletedFilter)

	filter := bson.M{
		"$or":  filterData1,
		"$and": filterData2,
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

	err = cur.All(ctx, &res)
	if err != nil {
		return
	}

	rows, _ = u.users.CountDocuments(ctx, filter)

	return
}
