package domain

import "context"

type PaymentUsecase interface {
	Pay(ctx context.Context, userID string, amount float64, remarks string) (*ResponsePayment, error)
}

type ResponsePayment struct {
	PaymentID     string  `json:"payment_id"`
	Amount        float64 `json:"amount"`
	Remarks       string  `json:"remarks"`
	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`
	CreatedDate   string  `json:"created_date"`
}
