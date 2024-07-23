package response

import (
	"paywatcher/src/domain/userdomain"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID       uuid.UUID         `json:"id"`
	Email    string            `json:"email"`
	Username string            `json:"username"`
	Role     []userdomain.Role `json:"role"`
}

func NewUserResponse(user *userdomain.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
	}
}
