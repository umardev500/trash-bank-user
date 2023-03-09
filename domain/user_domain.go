package domain

import (
	"context"
	"user/pb"
	"user/pb/user"

	"go.mongodb.org/mongo-driver/bson"
)

type UserService interface {
	Find(ctx context.Context, req *user.UserFindRequest) (res *user.UserFindResponse, err error)
	FindOne(ctx context.Context, req *user.UserFindOneRequest) (res *user.UserFindOneResponse, err error)
	Update(ctx context.Context, req *user.UserUpdateRequest) (res *pb.OperationResponse, err error)
}

type UserRepo interface {
	Find(ctx context.Context, sort, pages, perPage int64, status, search string) (res []User, rows int64, err error)
	FindOne(ctx context.Context, userId, user, pass string, isLogin bool) (res *User, err error)
	Update(ctx context.Context, userId string, payload bson.D) (res *pb.OperationResponse, err error)
}

type User struct {
	UserId    string        `bson:"user_id"`
	User      string        `bson:"user"`
	Pass      string        `bson:"pass"`
	Level     string        `bson:"level"`
	Details   *UserDetails  `bson:"details"`
	Status    []*UserStatus `bson:"status"`
	CreatedAt int64         `bson:"created_at"`
	UpdatedAt int64         `bson:"updated_at"`
}
