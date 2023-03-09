package repo

import (
	"context"
	"user/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (u *userRepo) Update(ctx context.Context, userId string, payload bson.D, status bson.M) (res *pb.OperationResponse, err error) {
	affected := true
	filter := bson.M{"user_id": userId}
	var resp *mongo.UpdateResult
	if status != nil {
		statusSet := bson.M{
			"status": bson.M{
				"$each":     []bson.M{status},
				"$position": 0,
			},
		}
		resp, err = u.users.UpdateOne(ctx, filter, bson.M{"$set": payload, "$push": statusSet})
	} else {
		resp, err = u.users.UpdateOne(ctx, filter, bson.M{"$set": payload})
	}
	if err != nil {
		return
	}

	affected = resp.ModifiedCount > 0

	res = &pb.OperationResponse{IsAffected: affected}

	return
}
