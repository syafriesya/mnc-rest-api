package login

import "mnc-rest-api/internal/domain"

type LoginUsecase struct {
	userRepo  domain.UserRepository
	jwtSecret string
}

func New(
	jwtSecret string,
	userRepo domain.UserRepository,
) LoginUsecase {
	return LoginUsecase{
		jwtSecret: jwtSecret,
		userRepo:  userRepo,
	}
}
