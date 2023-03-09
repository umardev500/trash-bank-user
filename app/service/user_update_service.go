package service

import (
	"context"
	"errors"
	"net/http"
	"user/helper"
	"user/pb"
	"user/pb/user"

	"go.mongodb.org/mongo-driver/bson"
)

func (u *userService) Update(ctx context.Context, req *user.UserUpdateRequest) (res *pb.OperationResponse, err error) {
	userId := req.UserId
	payload := req.Payload

	var userData bson.D
	if payload != nil {
		userData = bson.D{
			{Key: "user", Value: payload.User},
			{Key: "pass", Value: payload.Pass},
			{Key: "level", Value: payload.Level},
			{Key: "updated_at", Value: 1000},
		}

		helper.RemoveEmpty(userData, &userData)
	}

	if userData == nil {
		err = errors.New(http.StatusText(http.StatusBadRequest))
		return
	}

	// do update
	res, err = u.repo.Update(ctx, userId, userData)
	if err != nil {
		return
	}

	return
}
