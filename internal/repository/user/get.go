package user

import (
	"context"
	"errors"
	"mnc-rest-api/internal/domain"

	"gorm.io/gorm"
)

func (r UserRepository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Where("phone_number = ?", phoneNumber).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, err
		}
	}
	return user, nil
}

func (r UserRepository) GetUserByID(ctx context.Context, userID string) (domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}
