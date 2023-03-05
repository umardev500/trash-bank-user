package domain

type UserStatus struct {
	StatusText string `bson:"status_text"`
	StatusTime int64  `bson:"status_time"`
}
