package domain

import (
	"context"
	"time"
)

type TopupUsecase interface {
	TopUp(ctx context.Context, userID string, amount float64) (ResponseTopUp, error)
}

type TopupRepository interface {
	CreateTopUpRecord(ctx context.Context, topUp TopUp) error
}

type ResponseTopUp struct {
	TopUpID       string    `json:"top_up_id"`
	AmountTopUp   float64   `json:"amount_top_up"`
	BalanceBefore float64   `json:"balance_before"`
	BalanceAfter  float64   `json:"balance_after"`
	CreatedDate   time.Time `json:"created_date"`
}
