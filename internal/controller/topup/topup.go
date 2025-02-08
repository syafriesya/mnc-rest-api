package topup

import "mnc-rest-api/internal/domain"

type TopupController struct {
	topupUsecase domain.TopupUsecase
}

func New(topupUsecase domain.TopupUsecase) TopupController {
	return TopupController{
		topupUsecase: topupUsecase,
	}
}
