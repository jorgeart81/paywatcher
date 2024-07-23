package request

import (
	"fmt"
	"paywatcher/src/domain/userdomain"
)

type RegisterUser struct {
	Email    string   `form:"email" json:"email" binding:"email,required"`
	Username string   `form:"usename" json:"username" binding:"required"`
	Password string   `form:"password" json:"password,omitempty" binding:"required"`
	Role     []string `form:"role" json:"role"`
}

// ValidateRoles checks if all roles assigned to the user are allowed.
func (u *RegisterUser) ValidateRoles() error {
	for _, role := range u.Role {
		if _, ok := userdomain.AllowedRoles[role]; !ok {
			return fmt.Errorf("invalid role: %s", role)
		}
	}
	return nil
}

type LoginUser struct {
	Email    string `form:"email" json:"email" binding:"email,required"`
	Password string `form:"password" json:"password,omitempty" binding:"required"`
}
