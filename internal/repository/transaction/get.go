package transaction

import (
	"context"
	"mnc-rest-api/internal/domain"
)

func (r TransactionRepository) FindByUserID(ctx context.Context, userID string) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := r.db.Where("user_id = ?", userID).Order("created_date desc").Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
