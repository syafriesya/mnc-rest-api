package domain

import "time"

type TopUp struct {
	TopUpID       string    `json:"top_up_id" gorm:"column:top_up_id;primary_key"`
	UserID        string    `json:"user_id" gorm:"column:user_id"`
	Amount        float64   `json:"amount" gorm:"column:amount"`
	BalanceBefore float64   `json:"balance_before" gorm:"column:balance_before"`
	BalanceAfter  float64   `json:"balance_after" gorm:"column:balance_after"`
	CreatedDate   time.Time `json:"created_date" gorm:"column:created_date"`
}
