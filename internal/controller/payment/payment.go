package payment

import "mnc-rest-api/internal/domain"

type PaymentController struct {
	paymentUsecase domain.PaymentUsecase
}

func New(paymentUsecase domain.PaymentUsecase) PaymentController {
	return PaymentController{
		paymentUsecase: paymentUsecase,
	}
}
