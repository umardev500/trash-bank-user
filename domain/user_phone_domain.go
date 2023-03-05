package domain

type UserPhone struct {
	Number string `bson:"number"`
	IsWa   bool   `bson:"is_wa"`
}
