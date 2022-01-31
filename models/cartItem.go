package models

import (
	"time"
)

type (
    // AgeRatingCategory
    CartItem struct {
        ID          uint        `gorm:"primary_key" json:"id"`
        ProductID   uint		`json:"productID"`
		UserID		uint 		`json:"userID"`
		CartID		uint		`json:"cartID"`
        CreatedAt   time.Time   `json:"created_at"`
        UpdatedAt   time.Time   `json:"updated_at"`
        User        User   		`json:"-"`
		Cart		Cart		`json:"-"`
		Product		Product		`json:"-"`
    }
)