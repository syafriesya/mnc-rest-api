package topup

import (
	"context"
	"mnc-rest-api/internal/domain"
	"time"

	"github.com/google/uuid"
)

func (u TopupUsecase) TopUp(ctx context.Context, userID string, amount float64) (domain.ResponseTopUp, error) {

	user, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return domain.ResponseTopUp{}, err
	}

	balanceBefore := user.Balance

	user.Balance += amount

	if err := u.userRepo.UpdateUserBalance(ctx, user); err != nil {
		return domain.ResponseTopUp{}, err
	}

	topUp := domain.TopUp{
		TopUpID:       uuid.New().String(),
		UserID:        userID,
		Amount:        amount,
		BalanceBefore: balanceBefore,
		BalanceAfter:  user.Balance,
		CreatedDate:   time.Now(),
	}

	if err := u.topupRepo.CreateTopUpRecord(ctx, topUp); err != nil {
		return domain.ResponseTopUp{}, err
	}

	transaction := domain.Transaction{
		TransactionID:   uuid.New().String(),
		UserID:          userID,
		TopUpID:         topUp.TopUpID,
		TransactionType: "CREDIT",
		Amount:          amount,
		BalanceBefore:   balanceBefore,
		BalanceAfter:    user.Balance,
		Remarks:         "Top-up balance",
		CreatedDate:     time.Now(),
	}

	if err := u.transactionRepo.CreateTransactionRecord(ctx, transaction); err != nil {
		return domain.ResponseTopUp{}, err
	}

	return domain.ResponseTopUp{
		TopUpID:       topUp.TopUpID,
		AmountTopUp:   topUp.Amount,
		BalanceBefore: balanceBefore,
		BalanceAfter:  user.Balance,
		CreatedDate:   topUp.CreatedDate,
	}, nil
}
