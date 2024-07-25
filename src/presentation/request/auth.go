package request

import (
	"fmt"
	"paywatcher/src/domain/entity"
)

type RegisterUser struct {
	Email    string   `form:"email" json:"email" binding:"email,required"`
	Username string   `form:"usename" json:"username" binding:"required"`
	Password string   `form:"password" json:"password,omitempty" binding:"required"`
	Role     []string `form:"role" json:"role"`
}

func (u *RegisterUser) ToUserEntity() *entity.UserEnt {
	return &entity.UserEnt{
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
		Role:     u.Role,
	}
}

// ValidateRoles checks if all roles assigned to the user are allowed.
func (u *RegisterUser) ValidateRoles() error {
	for _, role := range u.Role {
		if _, ok := entity.UserAllowedRoles[role]; !ok {
			return fmt.Errorf("invalid role: %s", role)
		}
	}
	return nil
}

type LoginUser struct {
	Email    string `form:"email" json:"email" binding:"email,required"`
	Password string `form:"password" json:"password,omitempty" binding:"required"`
}
