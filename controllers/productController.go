package controllers

import (
	"net/http"
	"time"

	"fp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productInput struct {
    Name        		string `json:"name"`
    Price        		string `json:"price"`
    Image        		string `json:"image"`
    Description        	string `json:"description"`
    ProductCategoryID   uint `json:"product_category_id"`
    SellerID   			uint `json:"seller_id"`
}

// GetAllProduct godoc
// @Summary Get all Product.
// @Description Get a list of Product.
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Product
// @Router /product [get]
func GetAllProduct(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var products []models.ProductCategory
    db.Find(&products)

    c.JSON(http.StatusOK, gin.H{"data": products})
}

// Create a Product godoc
// @Summary Create Product 
// @Description create new Product
// @Tags Product
// @Param Body body productInput true "the body to create new product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Product
// @Router /product [post]
func CreateProduct(c *gin.Context) {
    var input productInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    product := models.Product{Name: input.Name,Price: input.Price,Image: input.Image,Description: input.Description,ProductCategoryID: input.ProductCategoryID, SellerID: input.SellerID}
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)

    db.Create(&product)
    c.JSON(http.StatusOK, gin.H{"data": product})
}
// GetProductById godoc
// @Summary Get Product.
// @Description Get an Productby id.
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Success 200 {object} models.Product
// @Router /product/{id} [get]
func GetProductById(c *gin.Context) { // Get model if exist
    var product models.Product

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": product})
}



// UpdateProduct godoc
// @Summary Update Product.
// @Description Update Product by id.
// @Tags Product
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Product id"
// @Param Body body productInput true "the body to update product"
// @Success 200 {object} models.Product
// @Router /product/{id} [patch]
func UpdateProduct(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var product models.Product
    if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input productInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Product
    updatedInput.Name = input.Name
	updatedInput.Price= input.Price
	updatedInput.Image= input.Image
	updatedInput.Description= input.Description
	updatedInput.ProductCategoryID= input.ProductCategoryID
	updatedInput.SellerID= input.SellerID
    updatedInput.UpdatedAt = time.Now()

    db.Model(&product).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct godoc
// @Summary Delete one Product.
// @Description Delete a Product by id.
// @Tags Product
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "ProductCategory id"
// @Success 200 {object} map[string]boolean
// @Router /product/{id} [delete]
func DeleteProduct(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var product models.Product
    if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&product)

    c.JSON(http.StatusOK, gin.H{"data": true})
}