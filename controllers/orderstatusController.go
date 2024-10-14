package controllers

import (
	"dathang/databases"
	"dathang/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create OrderStatus
// @Tags Order Status
// @Param orderstatus body models.OrderStatus true "OrderStatus data"
// @Router /orderstatus/ [post]
func CreateOrderStatus(c *gin.Context) {
	var input struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"mes": "Lỗi đầu vào json",
		})
		return
	}

	createdAt := time.Now()

	orderstatus := models.OrderStatus{
		Status:    input.Status,
		CreatedAt: &createdAt,
	}

	if err := databases.DB.Create(&orderstatus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không tạo được order status",
		})
		return
	}

	c.JSON(http.StatusCreated, orderstatus)

}

// @Summary Get ALL ShippingMethod
// @Tags Shipping Method
// @Router /orderstatus/ [get]
func GetAllOrderStatus(c *gin.Context) {
	var orderstatus []models.OrderStatus
	if err := databases.DB.Find(&orderstatus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi truy vấn order status",
		})
		return
	}

	c.JSON(http.StatusOK, orderstatus)
}
