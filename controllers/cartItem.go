package controllers

import (
	"net/http"
	"time"

	"fp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type cartItemInput struct {
    ProductID        		uint `json:"product_id"`
    UserID                  uint   `json:"user_id"`
}

// GetAllCartItem godoc
// @Summary Get all CartItem.
// @Description Get a list of CartItem.
// @Tags CartItem
// @Produce json
// @Success 200 {object} []models.CartItem
// @Router /cart-item [get]
func GetAllCartItem(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var cartItem []models.CartItem
    db.Find(&cartItem)

    c.JSON(http.StatusOK, gin.H{"data": cartItem})
}

// Create a CartItem godoc
// @Summary Create CartItem
// @Description create new CartItem
// @Tags CartItem
// @Param Body body cartItemInput true "the body to create new cartItem"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.CartItem
// @Router /cart-item [post]
func CreateCartItem(c *gin.Context) {
    var input cartItemInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    cartItem := models.CartItem{ProductID: input.ProductID,UserID:input.UserID}
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)

    db.Create(&cartItem)
    c.JSON(http.StatusOK, gin.H{"data": cartItem})
}
// GetCartItemByIdUser godoc
// @Summary Get CartItem.
// @Description Get an CartItem by id.
// @Tags CartItem
// @Produce json
// @Param id path string true "CartItem id"
// @Success 200 {object} models.CartItem
// @Router /cart-item/{id} [get]
func GetCartItemByIdUser(c *gin.Context) { // Get model if exist
    var cartItem models.CartItem

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("user_id = ?", c.Param("user_id")).Find(&cartItem).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": cartItem})
}

// UpdateCartItem godoc
// @Summary Update CartItem.
// @Description Update CartItem by id.
// @Tags CartItem
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "CartItem id"
// @Param Body body cartItemInput true "the body to update cartItem"
// @Success 200 {object} models.CartItem
// @Router /cart-item/{id} [patch]
func UpdateCartItem(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var cartItem models.CartItem
    if err := db.Where("id = ?", c.Param("id")).First(&cartItem).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input cartItemInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.CartItem
    updatedInput.ProductID = input.ProductID
    updatedInput.UserID = input.UserID
    updatedInput.UpdatedAt = time.Now()

    db.Model(&cartItem).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": cartItem})
}

// DeleteCartItem godoc
// @Summary Delete one CartItem.
// @Description Delete a CartItem by id.
// @Tags CartItem
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "CartItem id"
// @Success 200 {object} map[string]boolean
// @Router /cart-item/{id} [delete]
func DeleteCartItem(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var cartItem models.CartItem
    if err := db.Where("id = ?", c.Param("id")).First(&cartItem).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&cartItem)

    c.JSON(http.StatusOK, gin.H{"data": true})
}