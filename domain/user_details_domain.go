package domain

type UserDetails struct {
	Name    string       `bson:"name"`
	Email   string       `bson:"email"`
	Phone   *UserPhone   `bson:"phone"`
	Address *UserAddress `bson:"address"`
}
