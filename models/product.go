package models

import (
	"time"
)

type (
    // AgeRatingCategory
    Product struct {
        ID                  uint               `gorm:"primary_key" json:"id"`
        Name                string            `json:"name"`
        Price 		        string            `json:"price"`
        Image 		        string            `json:"image"`
        Description         string            `json:"description"`
        ProductCategoryID   uint                `json:"productCategoryID"`
        SellerID            uint                `json:"sellerID"`
        CreatedAt           time.Time         `json:"created_at"`
        UpdatedAt           time.Time         `json:"updated_at"`
        ProductCategory     ProductCategory   `json:"-"`
        Discussion          []Discussion      `json:"-"`
        CartItem            []CartItem        `json:"-"`
        Seller               Seller          `json:"-"`
    }
)