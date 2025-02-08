package register

import (
	"context"
	"mnc-rest-api/internal/domain"
)

func (r RegisterRepository) RegisterUser(ctx context.Context, user domain.User) error {
	return r.db.WithContext(ctx).Create(&user).Error
}
