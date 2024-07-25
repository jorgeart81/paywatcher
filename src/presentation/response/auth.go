package response

import (
	"paywatcher/src/domain/entity"

	"github.com/google/uuid"
)

type AuthResponse struct {
	ID       uuid.UUID   `json:"id"`
	Email    string      `json:"email"`
	Username string      `json:"username"`
	Role     []string    `json:"role"`
	Tokens   interface{} `json:"tokens"`
}

func NewAuthResponse(user *entity.UserEnt, tokens interface{}) AuthResponse {
	return AuthResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
		Tokens:   tokens,
	}
}
