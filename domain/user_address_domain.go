package domain

type UserAddress struct {
	Province   string `bson:"province"`
	City       string `bson:"city"`
	District   string `bson:"district"`
	Village    string `bson:"village"`
	PostalCode string `bson:"postal_code"`
	Detail     string `bson:"detail"`
}
