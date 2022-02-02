package controllers

import (
	"net/http"
	"time"

	"fp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sellerInput struct {
    StoreName        		string `json:"store_name"`
    UserID                  uint   `json:"user_id"`
}

// GetAllSeller godoc
// @Summary Get all Seller.
// @Description Get a list of Seller.
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Seller
// @Router /seller [get]
func GetAllSeleer(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var sellers []models.Seller
    db.Find(&sellers)

    c.JSON(http.StatusOK, gin.H{"data": sellers})
}

// Create a Seller godoc
// @Summary Create Seller 
// @Description create new Seller
// @Tags Seller
// @Param Body body sellerInput true "the body to create new seller"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Seller
// @Router /seller [post]
func CreateSeller(c *gin.Context) {
    var input sellerInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    seller := models.Seller{StoreName: input.StoreName,UserID:input.UserID}
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)

    db.Create(&seller)
    c.JSON(http.StatusOK, gin.H{"data": seller})
}
// GetSellerById godoc
// @Summary Get Seller.
// @Description Get an Seller by id.
// @Tags Product
// @Produce json
// @Param id path string true "Seller id"
// @Success 200 {object} models.Seller
// @Router /seller/{id} [get]
func GetSellerById(c *gin.Context) { // Get model if exist
    var seller models.Seller

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&seller).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": seller})
}



// UpdateSeller godoc
// @Summary Update Seller.
// @Description Update Seller by id.
// @Tags Seller
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Seller id"
// @Param Body body sellerInput true "the body to update seller"
// @Success 200 {object} models.Seller
// @Router /seller/{id} [patch]
func UpdateSeller(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var seller models.Seller
    if err := db.Where("id = ?", c.Param("id")).First(&seller).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input sellerInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Seller
    updatedInput.StoreName = input.StoreName
    updatedInput.UserID = input.UserID
    updatedInput.UpdatedAt = time.Now()

    db.Model(&seller).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": seller})
}

// DeleteSeller godoc
// @Summary Delete one Seller.
// @Description Delete a Seller by id.
// @Tags Seller
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Seller id"
// @Success 200 {object} map[string]boolean
// @Router /seller/{id} [delete]
func DeleteSeller(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var seller models.Seller
    if err := db.Where("id = ?", c.Param("id")).First(&seller).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&seller)

    c.JSON(http.StatusOK, gin.H{"data": true})
}