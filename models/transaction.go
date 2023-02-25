package models

import (
	"time"
)

// User struct

type Transaction struct {
	// gorm.Model
	ID          uint      `gorm:"primarykey"`
	Date        time.Time `json:"date" `
	Description string    `json:"description" gorm:"column:description"`
	Credit      int       `json:"credit" gorm:"column:credit"`
	Debit       int       `json:"debit" gorm:"column:debit"`
	Balance     int       `json:"balance" gorm:"column:balance"`
}
