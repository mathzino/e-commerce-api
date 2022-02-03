package models

import (
	"time"
)

type (
    // User
    Seller struct {
        ID          uint            `json:"id" gorm:"primary_key"`
        StoreName   string       	`json:"store_name"`
        UserID      uint            `json:"userID"`
        CreatedAt   time.Time       `json:"created_at"`
        UpdatedAt   time.Time       `json:"updated_at"`
        // Transaction []Transaction    `json:"-"`
		Product		[]Product		`json:"-"`
       
    }
)





