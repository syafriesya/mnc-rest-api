package topup

import (
	"context"
	"mnc-rest-api/internal/domain"
)

func (r TopupRepository) CreateTopUpRecord(ctx context.Context, topUp domain.TopUp) error {
	return r.db.WithContext(ctx).Create(&topUp).Error
}
