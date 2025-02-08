package topup

import "gorm.io/gorm"

type TopupRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) TopupRepository {
	return TopupRepository{
		db: db,
	}
}
