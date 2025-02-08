package register

import "mnc-rest-api/internal/domain"

type RegisterController struct {
	registerUsecase domain.RegisterUsecase
}

func New(registerUsecase domain.RegisterUsecase) RegisterController {
	return RegisterController{
		registerUsecase: registerUsecase,
	}
}
