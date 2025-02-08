package domain

import (
	"context"
	"time"
)

type Transaction struct {
	TransactionID   string    `json:"transaction_id" gorm:"primary_key"`
	UserID          string    `json:"user_id"`
	TopUpID         string    `json:"top_up_id,omitempty"`
	TransferID      string    `json:"transfer_id,omitempty"`
	PaymentID       string    `json:"payment_id,omitempty"`
	TransactionType string    `json:"transaction_type"`
	Amount          float64   `json:"amount"`
	BalanceBefore   float64   `json:"balance_before"`
	BalanceAfter    float64   `json:"balance_after"`
	Remarks         string    `json:"remarks"`
	CreatedDate     time.Time `json:"created_date"`
}

type TransactionRepository interface {
	FindByUserID(ctx context.Context, userID string) ([]Transaction, error)
	CreateTransactionRecord(ctx context.Context, transaction Transaction) error
}

type TransactionUsecase interface {
	GetUserTransactions(ctx context.Context, userID string) ([]Transaction, error)
}
