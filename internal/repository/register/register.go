package register

import "gorm.io/gorm"

type RegisterRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RegisterRepository {
	return RegisterRepository{
		db: db,
	}
}
