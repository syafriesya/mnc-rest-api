package user

import (
	"context"
	"fmt"
	"mnc-rest-api/internal/domain"
)

func (r UserRepository) UpdateUserBalance(ctx context.Context, user domain.User) error {
	return r.db.WithContext(ctx).Model(&user).Update("balance", user.Balance).Error
}

func (r UserRepository) UpdateUserProfile(ctx context.Context, user domain.User) error {
	if err := r.db.WithContext(ctx).Save(&user).Error; err != nil {
		return fmt.Errorf("failed to update user profile: %w", err)
	}
	return nil
}
