package payment

import (
	"context"
	"fmt"
	"mnc-rest-api/internal/domain"
	"time"

	"github.com/google/uuid"
)

func (p PaymentUsecase) Pay(ctx context.Context, userID string, amount float64, remarks string) (*domain.ResponsePayment, error) {
	user, err := p.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user.Balance < amount {
		return nil, fmt.Errorf("balance is not enough")
	}

	balanceBefore := user.Balance
	balanceAfter := balanceBefore - amount

	paymentID := uuid.New().String()
	user.Balance = balanceAfter
	if err := p.userRepo.UpdateUserBalance(ctx, user); err != nil {
		return nil, err
	}

	transaction := domain.Transaction{
		TransactionID:   uuid.New().String(),
		UserID:          userID,
		PaymentID:       paymentID,
		TransactionType: "DEBIT",
		Amount:          amount,
		BalanceBefore:   balanceBefore,
		BalanceAfter:    balanceAfter,
		Remarks:         remarks,
		CreatedDate:     time.Now(),
	}

	if err := p.transactionRepo.CreateTransactionRecord(ctx, transaction); err != nil {
		return nil, err
	}

	response := &domain.ResponsePayment{
		PaymentID:     paymentID,
		Amount:        amount,
		Remarks:       remarks,
		BalanceBefore: balanceBefore,
		BalanceAfter:  balanceAfter,
		CreatedDate:   time.Now().Format("2006-01-02 15:04:05"),
	}

	return response, nil
}
