package login

import "gorm.io/gorm"

type LoginRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) LoginRepository {
	return LoginRepository{
		db: db,
	}
}
