package models

import (
	"time"
)

type (
    // AgeRatingCategory
    Discussion struct {
        ID          int      	`gorm:"primary_key" json:"id"`
        Discussion string    	`json:"discussion"`
		UserID		uint     	`json:"userID"`
		ProductID	uint		`json:"productID"`
        CreatedAt   time.Time 	`json:"created_at"`
        UpdatedAt   time.Time 	`json:"updated_at"`
        User      	User   		`json:"-"`
		Product		Product		`json:"-"`
    }
)