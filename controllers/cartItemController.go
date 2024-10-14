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

// @Summary Get all cart item
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

// @Summary Get cart item id
// @Tags Cart Item
// @Param id path int true "ID"
// @Router /cartitem/{id} [get]
func GetCartItemById(c *gin.Context) {
	var cartitem models.CartItem
	if err := databases.DB.Preload("Product").Where("id = ?", c.Param("id")).First(&cartitem).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
			"mes": " Lỗi không tìm thấy id",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"mes":  "Cart item mà bạn muốn tìm:",
		"data": cartitem,
	})
}

// @Summary Delete cart item
// @Tags Cart Item
// @Param id path int true "ID"
// @Router /cartitem/{id} [delete]
func DeleteCartItem(c *gin.Context) {
	var cartitem models.CartItem
	if err := databases.DB.Where("id = ?", c.Param("id")).First(&cartitem).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
			"mes": " Lỗi không tìm thấy id",
		})
		return
	}

	if err := databases.DB.Delete(&cartitem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không xóa được.",
		})
		return
	}
	c.JSON(http.StatusOK, "Xóa thành công")
}

// @Summary Update cart item
// @Tags Cart Item
// @Param id path int true "ID"
// @Param cartitem body models.CartItem true "CartItem data"
// @Router /cartitem/{id} [put]
func UpdateCartItem(c *gin.Context) {
	var input struct {
		CartID    uint `json:"cart_id"`
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"mes": "Lỗi nhập không đúng định dạng json",
		})
		return
	}
	// Khởi tạo slice mới để lưu các cart_item vừa được cập nhật
	

	var cart models.Cart

	if err := databases.DB.Where("id =?", input.CartID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không tìm thấy cart",
		})
		return
	}

	var product models.Product
	if err := databases.DB.Where("id = ?", input.ProductID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không tìm thấy product",
		})
		return
	}

	var cartitem models.CartItem

	if err := databases.DB.Where("id = ?", c.Param("id")).First(&cartitem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không tìm thấy id cart item",
		})
		return
	}

	cartitem.CartID = input.CartID
	cartitem.ProductID = input.ProductID
	cartitem.Quantity = input.Quantity
	cartitem.Price = product.Price * float64(input.Quantity)

	if err := databases.DB.Save(&cartitem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không thể lưu vào data base",
		})
		return
	}
	c.JSON(http.StatusOK, cartitem)
}
