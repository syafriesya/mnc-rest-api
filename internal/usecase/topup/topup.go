package topup

import "mnc-rest-api/internal/domain"

type TopupUsecase struct {
	topupRepo       domain.TopupRepository
	userRepo        domain.UserRepository
	transactionRepo domain.TransactionRepository
}

func New(
	topupRepo domain.TopupRepository,
	userRepo domain.UserRepository,
	transactionRepo domain.TransactionRepository,
) TopupUsecase {
	return TopupUsecase{
		topupRepo:       topupRepo,
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
	}
}
