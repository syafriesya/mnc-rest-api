package user

import "mnc-rest-api/internal/domain"

type UserController struct {
	userUsecase domain.UserUsecase
}

func New(userUsecase domain.UserUsecase) UserController {
	return UserController{
		userUsecase: userUsecase,
	}
}
