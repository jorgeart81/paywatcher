package usecases

import (
	"errors"
	"paywatcher/src/application/auth"
	"paywatcher/src/domain/userdomain"

	"golang.org/x/crypto/bcrypt"
)

type LoginUserUseCase struct {
	userRepo userdomain.UserRepository
	auth     *auth.Auth
}

func NewLoginUserUseCase(userRepo userdomain.UserRepository, auth *auth.Auth) LoginUserUseCase {
	return LoginUserUseCase{
		userRepo: userRepo,
		auth:     auth,
	}
}

func (uc *LoginUserUseCase) Execute(email, password string) (*userdomain.User, string, error) {
	repo := uc.userRepo

	user, err := repo.GetUserByEmail(email)
	if err != nil {
		return nil, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	jwtUser := auth.JwtUser{
		ID:       user.ID,
		Username: user.Username,
	}

	tokenPairs, err := uc.auth.GenerateTokenPair(&jwtUser)
	if err != nil {
		return nil, "", err
	}

	return user, tokenPairs.AccessToken, nil
}
