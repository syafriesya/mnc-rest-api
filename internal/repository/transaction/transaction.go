package transaction

import "gorm.io/gorm"

type TransactionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) TransactionRepository {
	return TransactionRepository{
		db: db,
	}
}
