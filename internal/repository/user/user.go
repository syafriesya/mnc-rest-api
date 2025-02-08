package user

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}
