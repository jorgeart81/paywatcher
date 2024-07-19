package userdomain

import "paywatcher/database/model"

type UserRepository interface {
	GetUserByEmail(email string) (*model.User, error)
}
