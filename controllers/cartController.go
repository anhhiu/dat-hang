package controllers

import (
	"dathang/databases"
	"dathang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create cart
// @Tags Cart
// @Param cart body models.Cart true "Cart data"
// @Router /cart/ [post]
func CreateCart(c *gin.Context) {
	var input struct {
		UserID uint `json:"user_id"`

		CartItems []struct {
			ProductID uint    `json:"product_id"`
			Quantity  int     `json:"quantity"`
			Price     float64 `json:"price"`
		} `json:"cart_items"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var tongTienCacDon float64
	cart := models.Cart{
		UserID: input.UserID,
	}

	if err := databases.DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	for _, item := range input.CartItems {
		var prd models.Product
		if err := databases.DB.Where("id = ?", item.ProductID).First(&prd); err != nil {
			c.JSON(500, err.Error)
			return
		}
		cartitem := models.CartItem{
			CartID:    cart.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     float64(item.Quantity) * prd.Price,
		}

		if err := databases.DB.Create(&cartitem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		//c.JSON(http.StatusCreated, cartitem)

		tongTienCacDon += cartitem.Price
	}

	cart.TotalPrice = tongTienCacDon

	if err := databases.DB.Preload("User").Save(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":    "Giỏ hàng đã được tạo thành công",
		"cart":       cart,
		"totalPrice": cart.TotalPrice,
	})
}

// @Summary Get all cart
// @Tags Cart
// @Router /cart/ [get]
func GetAllCart(c *gin.Context) {
	var carts []models.Cart

	if err := databases.DB.Preload("User").Find(&carts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, carts)
}
