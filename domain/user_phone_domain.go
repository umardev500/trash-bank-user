package domain

type UserPhone struct {
	Number string `bson:"number"`
	IsWa   string `bson:"is_wa"`
}
