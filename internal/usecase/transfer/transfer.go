package transfer

import "mnc-rest-api/internal/domain"

type TransferUsecase struct {
	userRepo        domain.UserRepository
	transactionRepo domain.TransactionRepository
}

func New(
	userRepo domain.UserRepository,
	transactionRepo domain.TransactionRepository,
) TransferUsecase {
	return TransferUsecase{
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
	}
}
