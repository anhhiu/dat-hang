package controllers

import (
	"dathang/databases"
	"dathang/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create review
// @Tags Review
// @Param review body models.Review true "Review data"
// @Router /review/ [post]
func CreateReview(c *gin.Context) {
	var input struct {
		ProductID int    `json:"product_id"`
		UserID    int    `json:"user_id"`
		Rating    int    `json:"rating"` // 1-5 sao
		Comment   string `json:"comment"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	sao := input.Rating
	if sao < 1 || sao > 5 {
		c.JSON(http.StatusBadRequest, "Enter the number 1 to 5")
		return
	}
	now := time.Now()
	review := models.Review{
		ProductID: input.ProductID,
		UserID:    input.UserID,
		Rating:    sao,
		Comment:   input.Comment,
		CreatedAt: &now,
	}

	if err := databases.DB.Preload("User").Preload("Product").
		Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, review)

}

// @Summary Get all review
// @Tags Review
// @Router /review/ [get]
func GetAllReview(c *gin.Context) {
	var reviews []models.Review

	if err := databases.DB.Preload("User").Preload("Product").
		Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// @Summary Get review by id
// @Tags Review
// @Param id path int true "ID"
// @Router /review/{id} [get]
func GetReviewById(c *gin.Context) {
	var review models.Review

	if err := databases.DB.Preload("User").Preload("Product").
		Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, review)
}

// @Summary Update review by id
// @Tags Review
// @Param id path int true "ID"
// @Param review body models.Review true "Review data"
// @Router /review/{id} [put]
func UpdateReview(c *gin.Context) {
	var input struct {
		ProductID int    `json:"product_id"`
		UserID    int    `json:"user_id"`
		Rating    int    `json:"rating"` // 1-5 sao
		Comment   string `json:"comment"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	sao := input.Rating
	if sao < 1 || sao > 5 {
		fmt.Println("sao :", sao)
		return
	}
	fmt.Println("sao :", sao)
	now := time.Now()

	var review models.Review

	if err := databases.DB.Preload("User").Preload("Product").
		Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	review.ProductID = input.ProductID
	review.UserID = input.UserID
	review.Rating = sao
	review.UpdatedAt = &now
	review.Comment = input.Comment

	if err := databases.DB.Preload("User").Preload("Product").
		Save(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, review)
}

// @Summary Delete review by id
// @Tags Review
// @Param id path int true "ID"
// @Router /review/{id} [delete]
func DeleteReview(c *gin.Context) {
	var review models.Review

	if err := databases.DB.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := databases.DB.Delete(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, true)
}
