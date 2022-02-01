package controllers

import (
	"net/http"
	"time"

	"fp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productCategoryInput struct {
    Name        string `json:"name"`
}

// GetAllProductCategory godoc
// @Summary Get all Product Category.
// @Description Get a list of ProductCategory.
// @Tags ProductCategory
// @Produce json
// @Success 200 {object} []models.ProductCategory
// @Router /product-categories [get]
func GetAllProductCategory(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var categories []models.ProductCategory
    db.Find(&categories)

    c.JSON(http.StatusOK, gin.H{"data": categories})
}

// Create a Category godoc
// @Summary Create Product Category
// @Description create new Product Category
// @Tags ProductCategory
// @Param Body body productCategoryInput true "the body to create new category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.ProductCategory
// @Router /product-categories [post]
func CreateCategory(c *gin.Context) {
    var input productCategoryInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    category := models.ProductCategory{Name: input.Name}
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)

    db.Create(&category)
    c.JSON(http.StatusOK, gin.H{"data": category})
}
// GetProductCategoryById godoc
// @Summary Get ProductCategory.
// @Description Get an ProductCategory by id.
// @Tags ProductCategory
// @Produce json
// @Param id path string true "ProductCategory id"
// @Success 200 {object} models.ProductCategory
// @Router /product-categories/{id} [get]
func GetCategoryById(c *gin.Context) { // Get model if exist
    var category models.ProductCategory

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": category})
}



// UpdateProductCategory godoc
// @Summary Update ProductCategory.
// @Description Update ProductCategory by id.
// @Tags ProductCategory
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "ProductCategory id"
// @Param Body body productCategoryInput true "the body to update product category"
// @Success 200 {object} models.ProductCategory
// @Router /product-categories/{id} [patch]
func UpdateCategory(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var category models.ProductCategory
    if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input productCategoryInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.ProductCategory
    updatedInput.Name = input.Name
    updatedInput.UpdatedAt = time.Now()

    db.Model(&category).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": category})
}

// DeleteProductCategory godoc
// @Summary Delete one ProductCategory.
// @Description Delete a ProductCategory by id.
// @Tags ProductCategory
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "ProductCategory id"
// @Success 200 {object} map[string]boolean
// @Router /product-categories/{id} [delete]
func DeleteProductCategory(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var category models.ProductCategory
    if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&category)

    c.JSON(http.StatusOK, gin.H{"data": true})
}