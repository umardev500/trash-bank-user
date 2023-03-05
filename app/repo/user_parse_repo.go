package repo

import (
	"user/domain"
	"user/pb/user"
)

func (u *userRepo) parseUser(each domain.User) *user.User {
	detail := each.Details

	var userDetails *user.UserDetails
	if detail != nil {
		// user phone
		var userPhone *user.UserPhone
		phone := detail.Phone
		if phone != nil {
			userPhone = &user.UserPhone{
				Number: phone.Number,
				IsWa:   phone.IsWa,
			}
		}

		// user address
		var userAddress *user.UserAddress
		address := detail.Address
		if address != nil {
			userAddress = &user.UserAddress{
				Province:   address.Province,
				City:       address.City,
				District:   address.District,
				Village:    address.Village,
				PostalCode: address.PostalCode,
				Detail:     address.Detail,
			}
		}

		// user details
		userDetails = &user.UserDetails{
			Name:    detail.Name,
			Email:   detail.Email,
			Phone:   userPhone,
			Address: userAddress,
		}
	}

	var userStatus []*user.UserStatus
	if len(each.Status) > 0 {
		for _, val := range each.Status {
			status := &user.UserStatus{
				StatusText: val.StatusText,
				StatusTime: val.StatusTime,
			}
			userStatus = append(userStatus, status)
		}
	}

	item := &user.User{
		UserId:    each.UserId,
		User:      each.User,
		Pass:      each.Pass,
		Level:     each.Level,
		Details:   userDetails,
		Status:    userStatus,
		CreatedAt: each.CreatedAt,
		UpdatedAt: each.UpdatedAt,
	}

	return item
}
