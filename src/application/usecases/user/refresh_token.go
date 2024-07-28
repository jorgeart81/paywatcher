package user

import (
	"paywatcher/src/domain/repositories"
	"paywatcher/src/domain/services"
)

type RefreshTokenUseCase struct {
	userRepo repositories.UserRepository
	auth     services.Authenticator
}

func NewRefreshTokenUseCase(userRepo repositories.UserRepository, auth services.Authenticator) RefreshTokenUseCase {
	return RefreshTokenUseCase{
		userRepo: userRepo,
		auth:     auth,
	}
}

func (r *RefreshTokenUseCase) Execute(refreshToken string) (services.TokenPairs, error) {
	auth := r.auth

	claims, err := auth.VerifyToken(refreshToken)
	if err != nil {
		return services.TokenPairs{}, err
	}

	user, err := r.userRepo.GetUserById(claims.ID)
	if err != nil {
		return services.TokenPairs{}, err
	}

	jwtUser := services.AuthUser{
		ID:       user.ID,
		Username: claims.Username,
	}

	tokenPairs, err := auth.GenerateTokenPair(&jwtUser)
	if err != nil {
		return services.TokenPairs{}, err
	}

	return tokenPairs, nil
}
