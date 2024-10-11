package controllers

import (
	"dathang/databases"
	"dathang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create cart item
// @Tags Cart Item
// @Param cartitem body models.CartItem true "CartItem data"
// @Router /cartitem/ [post]
func CreateCartItem(c *gin.Context) {
	var input struct {
		CartID    uint `json:"cart_id"`
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var cart models.Cart

	if err := databases.DB.Where("id =?", input.CartID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	var product models.Product
	if err := databases.DB.Where("id = ?", input.ProductID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	cartItem := models.CartItem{
		CartID:    input.CartID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Price:     product.Price * float64(input.Quantity),
	}

	if err := databases.DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, cartItem)

}

// @Summary Create cart item
// @Tags Cart Item
// @Router /cartitem/ [get]
func GetAllCartItem(c *gin.Context) {
	var cartitems []models.CartItem

	if err := databases.DB.Preload("Product").Find(&cartitems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cartitems)
}
