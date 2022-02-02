package controllers

import (
	"net/http"
	"time"

	"fp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type discussionInput struct {
	DiscussionValue string    	`json:"discussion_value"`
	UserID		uint     	`json:"user_id"`
	ProductID	uint		`json:"product_id"`
}

// GetAllDiscussion godoc
// @Summary Get all Discussion.
// @Description Get a list of Discussion.
// @Tags Discussion
// @Produce json
// @Success 200 {object} []models.Discussion
// @Router /discussion [get]
func GetAllDiscussion(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var discussions []models.Discussion
    db.Find(&discussions)

    c.JSON(http.StatusOK, gin.H{"data": discussions})
}

// Create a Discussion godoc
// @Summary Create Discussion 
// @Description create new Discussion
// @Tags Discussion
// @Param Body body discussionInput true "the body to create new discussion"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Discussion
// @Router /discussion [post]
func CreateDiscussion(c *gin.Context) {
    var input discussionInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    discussion := models.Discussion{DiscussionValue: input.DiscussionValue,UserID:input.UserID,ProductID: input.ProductID}
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)

    db.Create(&discussion)
    c.JSON(http.StatusOK, gin.H{"data": discussion})
}
// GetDiscussionById godoc
// @Summary Get Discussion
// @Description Get an Discussion by id.
// @Tags Discussion
// @Produce json
// @Param id path string true "Discussion id"
// @Success 200 {object} models.Discussion
// @Router /discussion/{id} [get]
func GetDiscussionById(c *gin.Context) { // Get model if exist
    var discussion models.Discussion

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&discussion).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": discussion})
}



// UpdateDiscussion godoc
// @Summary Update Discussion.
// @Description Update Discussion by id.
// @Tags Discussion
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Discussion id"
// @Param Body body discussionInput true "the body to update discussion"
// @Success 200 {object} models.Discussion
// @Router /discussion/{id} [patch]
func UpdateDiscussion(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var discussion models.Discussion
    if err := db.Where("id = ?", c.Param("id")).First(&discussion).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input discussionInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Discussion
    updatedInput.DiscussionValue = input.DiscussionValue
    updatedInput.UserID = input.UserID
    updatedInput.ProductID = input.ProductID
    updatedInput.UpdatedAt = time.Now()

    db.Model(&discussion).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": discussion})
}

// DeleteDiscussion godoc
// @Summary Delete one Discussion.
// @Description Delete a Discussion by id.
// @Tags Discussion
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Discussion id"
// @Success 200 {object} map[string]boolean
// @Router /discussion/{id} [delete]
func DeleteDiscussion(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var discussion models.Discussion
    if err := db.Where("id = ?", c.Param("id")).First(&discussion).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&discussion)

    c.JSON(http.StatusOK, gin.H{"data": true})
}