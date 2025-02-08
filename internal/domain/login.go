package domain

import "context"

type LoginUsecase interface {
	LoginUser(ctx context.Context, phoneNumber string, pin string) (string, string, error)
}
