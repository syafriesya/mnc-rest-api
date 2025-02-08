package register

import (
	"context"
	"errors"
	"mnc-rest-api/internal/domain"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (u RegisterUsecase) RegisterUser(ctx context.Context, request domain.User) (domain.ResponseRegisterUser, error) {
	existingUser, err := u.userRepo.GetUserByPhoneNumber(ctx, request.PhoneNumber)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			request.UserID = uuid.New().String()

			hashedPin, err := bcrypt.GenerateFromPassword([]byte(request.Pin), bcrypt.DefaultCost)
			if err != nil {
				return domain.ResponseRegisterUser{}, errors.New("failed to hash pin")
			}
			request.Pin = string(hashedPin)
			request.CreatedDate = time.Now()

			if err := u.registerRepo.RegisterUser(ctx, request); err != nil {
				return domain.ResponseRegisterUser{}, err
			}

			return domain.ResponseRegisterUser{
				UserID:      request.UserID,
				FirstName:   request.FirstName,
				LastName:    request.LastName,
				PhoneNumber: request.PhoneNumber,
				Address:     request.Address,
				CreatedDate: request.CreatedDate,
			}, nil
		}

		return domain.ResponseRegisterUser{}, err
	}

	if existingUser.PhoneNumber != "" {
		return domain.ResponseRegisterUser{}, errors.New("phone number already registered")
	}

	return domain.ResponseRegisterUser{}, nil
}
