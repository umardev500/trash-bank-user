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

	var detailData bson.D
	if payload.Details != nil {
		details := payload.Details
		phone := details.Phone
		var userPhone bson.D
		if phone != nil {
			userPhone = bson.D{
				{Key: "details.phone.number", Value: phone.Number},
				{Key: "details.phone.is_wa", Value: phone.IsWa},
			}
		}

		address := details.Address
		var userAddress bson.D
		if address != nil {
			userAddress = bson.D{
				{Key: "details.address.province", Value: address.Province},
				{Key: "details.address.city", Value: address.City},
				{Key: "details.address.district", Value: address.District},
				{Key: "details.address.village", Value: address.Village},
				{Key: "details.address.postal_code", Value: address.PostalCode},
				{Key: "details.address.detail", Value: address.Detail},
			}
		}

		detailData = bson.D{
			{Key: "details.name", Value: details.Name},
			{Key: "details.email", Value: details.Name},
		}
		detailData = append(detailData, userPhone...)
		detailData = append(detailData, userAddress...)
	}

	helper.RemoveEmpty(detailData, &detailData)

	updatedTime := helper.GetTime(nil)
	var userData bson.D
	if payload != nil {
		userData = bson.D{
			{Key: "user", Value: payload.User},
			{Key: "pass", Value: payload.Pass},
			{Key: "level", Value: payload.Level},
			{Key: "updated_at", Value: updatedTime},
		}

		helper.RemoveEmpty(userData, &userData)
	}
	userData = append(userData, detailData...)

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
