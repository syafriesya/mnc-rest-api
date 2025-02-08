package user

import (
	"context"
	"mnc-rest-api/internal/domain"
	"time"
)

func (u UserUsecase) UpdateProfile(ctx context.Context, userID string, req domain.UpdateProfileRequest) (domain.ResponseUpdateProfile, error) {

	user, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return domain.ResponseUpdateProfile{}, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Address = req.Address

	if err := u.userRepo.UpdateUserProfile(ctx, user); err != nil {
		return domain.ResponseUpdateProfile{}, err
	}

	return domain.ResponseUpdateProfile{
		UserID:      user.UserID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Address:     user.Address,
		UpdatedDate: time.Now(),
	}, nil
}
