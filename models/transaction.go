package models

import (
	"time"
)

type (
    // AgeRatingCategory
    Transaction struct {
        ID          uint      	`gorm:"primary_key" json:"id"`
		TotalPrice 	string    	`json:"total_price"`
		Status		bool		`json:"status"`
        CartID		uint		`json:"cartID"`
		UserID		uint		`json:"userID"`
        CreatedAt   time.Time 	`json:"created_at"`
        UpdatedAt   time.Time 	`json:"updated_at"`
        Cart      	Cart   		`json:"-"`
		User		User		`json:"-"`
    }
)