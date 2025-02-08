package login

import "mnc-rest-api/internal/domain"

type LoginController struct {
	loginUsecase domain.LoginUsecase
}

func New(loginUsecase domain.LoginUsecase) LoginController {
	return LoginController{
		loginUsecase: loginUsecase,
	}
}
