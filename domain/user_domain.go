package domain

import (
	"context"
	"user/pb/user"
)

type UserService interface {
	Find(ctx context.Context, req *user.UserFindRequest) (res *user.UserFindResponse, err error)
	FindOne(ctx context.Context, req *user.UserFindOneRequest) (res *user.UserFindOneResponse, err error)
}

type UserRepo interface {
	Find(ctx context.Context, sort, pages, perPage int64, status, search string) (res []User, rows int64, err error)
	FindOne(ctx context.Context, userId, user, pass string, isLogin bool) (res *User, err error)
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
