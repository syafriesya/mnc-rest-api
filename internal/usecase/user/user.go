package user

import "mnc-rest-api/internal/domain"

type UserUsecase struct {
	userRepo domain.UserRepository
}

func New(
	userRepo domain.UserRepository,
) UserUsecase {
	return UserUsecase{
		userRepo: userRepo,
	}
}
