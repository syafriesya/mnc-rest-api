package domain

import (
	"context"
	"time"
)

type UserRepository interface {
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (User, error)
	GetUserByID(ctx context.Context, userID string) (User, error)
	UpdateUserBalance(ctx context.Context, user User) error
	UpdateUserProfile(ctx context.Context, user User) error
}

type UserUsecase interface {
	UpdateProfile(ctx context.Context, userID string, req UpdateProfileRequest) (ResponseUpdateProfile, error)
}

type User struct {
	UserID      string    `json:"user_id" gorm:"column:user_id;primary_key"`
	FirstName   string    `json:"first_name" gorm:"column:first_name"`
	LastName    string    `json:"last_name" gorm:"column:last_name"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;unique"`
	Address     string    `json:"address" gorm:"column:address"`
	Pin         string    `json:"pin" gorm:"column:pin"`
	Balance     float64   `json:"balance" gorm:"column:balance;default:0"`
	CreatedDate time.Time `json:"created_date" gorm:"column:created_date"`
}

type UpdateProfileRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}

type ResponseUpdateProfile struct {
	UserID      string    `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Address     string    `json:"address"`
	UpdatedDate time.Time `json:"updated_date"`
}
