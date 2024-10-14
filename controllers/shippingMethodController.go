package controllers

import (
	"dathang/databases"
	"dathang/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create shippingmethod
// @Tags Shipping Method
// @Param shippingmethod body models.ShippingMethod true "ShippingMethod data"
// @Router /shippingmethod/ [post]
func CreateShippingMethod(c *gin.Context) {
	var input struct {
		Name        string `json:"name"`
		Discription string `json:"discription"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"mes": "Lỗi đầu vào json",
		})
		return
	}

	createdAt := time.Now()

	shippingmethod := models.ShippingMethod{
		Name:        input.Name,
		Description: input.Discription,

		CreatedAt: &createdAt,
	}

	if err := databases.DB.Create(&shippingmethod).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không tạo được shippingmethod",
		})
		return
	}

	c.JSON(http.StatusCreated, shippingmethod)

}

// @Summary Get ALL ShippingMethod
// @Tags Shipping Method
// @Router /shippingmethod/ [get]
func GetAllShippingMethod(c *gin.Context) {
	var shippingmethods []models.ShippingMethod
	if err := databases.DB.Find(&shippingmethods).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi truy vấn shippingmethod",
		})
		return
	}

	c.JSON(http.StatusOK, shippingmethods)
}
