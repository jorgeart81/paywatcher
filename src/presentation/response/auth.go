package response

import (
	"paywatcher/src/domain/entity"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Role     []string  `json:"role"`
}

func NewUserResponse(user *entity.UserEnt) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
	}
}
