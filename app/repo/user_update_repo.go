package repo

import (
	"context"
	"user/pb"

	"go.mongodb.org/mongo-driver/bson"
)

func (u *userRepo) Update(ctx context.Context, userId string, payload bson.D) (res *pb.OperationResponse, err error) {
	affected := true
	filter := bson.M{"user_id": userId}
	resp, err := u.users.UpdateOne(ctx, filter, bson.M{"$set": payload})
	if err != nil {
		return
	}

	affected = resp.ModifiedCount > 0

	res = &pb.OperationResponse{IsAffected: affected}

	return
}
