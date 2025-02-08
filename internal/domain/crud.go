package domain

import (
	"context"

	"gorm.io/gorm"
)

type CrudRepository interface {
	GetAll(ctx context.Context) ([]Items, error)
	GetById(ctx context.Context, id string) (Items, error)
	PostItem(ctx context.Context, request Items) error
	PutItem(ctx context.Context, id string, request Items) error
	DeleteItem(ctx context.Context, id string) error
}
type CrudUsecase interface {
	GetAll(ctx context.Context) ([]Items, error)
	GetById(ctx context.Context, id string) (Items, error)
	PostItem(ctx context.Context, request Items) error
	PutItem(ctx context.Context, id string, request Items) error
	DeleteItem(ctx context.Context, id string) error
}

type Items struct {
	gorm.Model
	Detail string `json:"detail" gorm:"column:detail;type:text"`
}
