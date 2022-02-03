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
        CreatedAt   time.Time   `json:"created_at"`
        UpdatedAt   time.Time   `json:"updated_at"`
        User        User   		`json:"-"`
		Product		Product		`json:"-"`
		// CartID		uint		`json:"cartID"`
		// Cart		Cart		`json:"-"`
    }
)