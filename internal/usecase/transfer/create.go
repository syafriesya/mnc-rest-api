package transfer

import (
	"context"
	"fmt"
	"mnc-rest-api/internal/domain"
	"time"

	"github.com/google/uuid"
)

func (t TransferUsecase) Transfer(ctx context.Context, senderID, targetUserID string, amount float64, remarks string) (*domain.ResponseTransfer, error) {
	sender, err := t.userRepo.GetUserByID(ctx, senderID)
	if err != nil {
		return nil, err
	}

	targetUser, err := t.userRepo.GetUserByID(ctx, targetUserID)
	if err != nil {
		return nil, fmt.Errorf("Target user not found")
	}

	if sender.Balance < amount {
		return nil, fmt.Errorf("Balance is not enough")
	}

	balanceBefore := sender.Balance
	balanceAfter := balanceBefore - amount
	sender.Balance = balanceAfter

	if err := t.userRepo.UpdateUserBalance(ctx, sender); err != nil {
		return nil, err
	}

	targetUser.Balance += amount

	if err := t.userRepo.UpdateUserBalance(ctx, targetUser); err != nil {
		return nil, err
	}

	transferID := uuid.New().String()

	transaction := domain.Transaction{
		TransactionID:   uuid.New().String(),
		UserID:          senderID,
		TransferID:      transferID,
		TransactionType: "DEBIT",
		Amount:          amount,
		BalanceBefore:   balanceBefore,
		BalanceAfter:    balanceAfter,
		Remarks:         remarks,
		CreatedDate:     time.Now(),
	}

	if err := t.transactionRepo.CreateTransactionRecord(ctx, transaction); err != nil {
		return nil, err
	}

	response := &domain.ResponseTransfer{
		TransferID:    transferID,
		Amount:        amount,
		Remarks:       remarks,
		BalanceBefore: balanceBefore,
		BalanceAfter:  balanceAfter,
		CreatedDate:   time.Now().Format("2006-01-02 15:04:05"),
	}

	return response, nil
}
