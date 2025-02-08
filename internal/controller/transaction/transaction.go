package transaction

import "mnc-rest-api/internal/domain"

type TransactionController struct {
	transactionUsecase domain.TransactionUsecase
}

func New(transactionUsecase domain.TransactionUsecase) TransactionController {
	return TransactionController{
		transactionUsecase: transactionUsecase,
	}
}
