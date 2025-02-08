package payment

import "mnc-rest-api/internal/domain"

type PaymentUsecase struct {
	userRepo        domain.UserRepository
	transactionRepo domain.TransactionRepository
}

func New(
	userRepo domain.UserRepository,
	transactionRepo domain.TransactionRepository,
) PaymentUsecase {
	return PaymentUsecase{
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
	}
}
