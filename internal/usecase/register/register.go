package register

import (
	"mnc-rest-api/internal/domain"
)

type RegisterUsecase struct {
	registerRepo domain.RegisterRepository
	userRepo     domain.UserRepository
}

func New(
	registerRepo domain.RegisterRepository,
	userRepo domain.UserRepository,
) RegisterUsecase {
	return RegisterUsecase{
		registerRepo: registerRepo,
		userRepo:     userRepo,
	}
}
