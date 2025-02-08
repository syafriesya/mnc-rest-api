package login

import (
	"context"
	"errors"
	"fmt"
	"mnc-rest-api/internal/domain"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (u LoginUsecase) LoginUser(ctx context.Context, phoneNumber string, pin string) (string, string, error) {
	user, err := u.userRepo.GetUserByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", errors.New("phone number not registered")
		}
		return "", "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pin), []byte(pin)); err != nil {
		return "", "", errors.New("incorrect pin")
	}

	accessToken, err := u.generateJWTToken(user, time.Minute*15)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := u.generateJWTToken(user, time.Hour*24*7)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Function to generate JWT token
func (u LoginUsecase) generateJWTToken(user domain.User, expiry time.Duration) (string, error) {
	fmt.Println("login userID", user.UserID)
	claims := jwt.MapClaims{
		"user_id": user.UserID,
		"phone":   user.PhoneNumber,
		"exp":     time.Now().Add(expiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.jwtSecret))
}
