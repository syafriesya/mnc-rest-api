package transaction

import (
	"context"
	"mnc-rest-api/internal/domain"
)

func (t TransactionUsecase) GetUserTransactions(ctx context.Context, userID string) ([]domain.Transaction, error) {
	transactions, err := t.transactionRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
