package transaction

import "mnc-rest-api/internal/domain"

type TransactionUsecase struct {
	transactionRepo domain.TransactionRepository
}

func New(
	transactionRepo domain.TransactionRepository,
) TransactionUsecase {
	return TransactionUsecase{
		transactionRepo: transactionRepo,
	}
}
