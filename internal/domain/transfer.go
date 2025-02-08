package domain

import "context"

type ResponseTransfer struct {
	TransferID    string  `json:"transfer_id"`
	Amount        float64 `json:"amount"`
	Remarks       string  `json:"remarks"`
	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`
	CreatedDate   string  `json:"created_date"`
}

type TransferUsecase interface {
	Transfer(ctx context.Context, senderID, targetUserID string, amount float64, remarks string) (*ResponseTransfer, error)
}
