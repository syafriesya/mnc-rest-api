package transaction

import (
	"context"
	"fmt"
	"mnc-rest-api/internal/domain"
)

func (r TransactionRepository) CreateTransactionRecord(ctx context.Context, transaction domain.Transaction) error {
	if err := r.db.WithContext(ctx).Create(&transaction).Error; err != nil {
		return fmt.Errorf("failed to create transaction record: %w", err)
	}
	return nil
}
