package domain

import (
	"context"
	"time"
)

type RegisterRepository interface {
	RegisterUser(ctx context.Context, user User) error
}

type RegisterUsecase interface {
	RegisterUser(ctx context.Context, request User) (ResponseRegisterUser, error)
}

type ResponseRegisterUser struct {
	UserID      string    `json:"userId" gorm:"column:user_id;type:string;primary_key"`
	FirstName   string    `json:"firstName" gorm:"column:first_name;type:string"`
	LastName    string    `json:"lastName"  gorm:"column:last_name;type:string"`
	PhoneNumber string    `json:"phoneNumber" gorm:"column:phone_number;type:string;unique"`
	Address     string    `json:"address" gorm:"column:address;type:string"`
	CreatedDate time.Time `json:"created_date" gorm:"created_at"`
}
