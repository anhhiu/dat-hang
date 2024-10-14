package controllers

import (
	"dathang/databases"
	"dathang/models"
	"fmt"
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
		var product models.Product
		if err := databases.DB.Where("id = ?", item.ProductID).First(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		cartitem := models.CartItem{
			CartID:    cart.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     float64(item.Quantity) * product.Price,
		}

		if err := databases.DB.Create(&cartitem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		//c.JSON(http.StatusCreated, cartitem)

		tongTienCacDon += cartitem.Price
		fmt.Printf("%f", tongTienCacDon)
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
	// Khởi tạo slice để chứa cart items
	var allCartItems []models.CartItem

	// Lặp qua từng giỏ hàng để lấy cart items
	for _, cart := range carts {
		var cartItems []models.CartItem
		if err := databases.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		allCartItems = append(allCartItems, cartItems...) // Thêm cart items vào slice chung
	}

	// Tạo cấu trúc phản hồi
	response := gin.H{
		"carts": carts,
		"data":  allCartItems,
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get cart by id
// @Tags Cart
// @Param id path int true "ID"
// @Router /cart/{id} [get]
func GetCartById(c *gin.Context) {
	var cart models.Cart
	if err := databases.DB.Preload("User").Where("id = ?", c.Param("id")).First(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cart)
}

// @Summary Delete cart
// @Tags Cart
// @Param id path int true "ID"
// @Router /cart/{id} [delete]
func DeleteCart(c *gin.Context) {
	var cart models.Cart
	if err := databases.DB.Where("id = ?", c.Param("id")).First(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Không tìm thấy id cart",
		})
		return
	}

	if err := databases.DB.Delete(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Không xóa được.",
		})
	}

	c.JSON(http.StatusOK, "Xóa Thành công")
}

// @Summary Update cart
// @Tags Cart
// @Param id path int true "ID"
// @Param cart body models.Cart true "Cart data"
// @Router /cart/{id} [put]
func UpdateCart(c *gin.Context) {
	var input struct {
		UserID    uint `json:"user_id"`
		CartItems []struct {
			ProductID uint    `json:"product_id"`
			Quantity  int     `json:"quantity"`
			Price     float64 `json:"price"`
		} `json:"cart_items"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"mes": "Lỗi nhập không đúng định dạng",
		})
		return
	}
	fmt.Printf("Input nhận được: UserID: %d\n", input.UserID)

	for i, item := range input.CartItems {
		fmt.Printf("CartItem %d - ProductID: %d, Quantity: %d, Price: %.2f\n", i+1, item.ProductID, item.Quantity, item.Price)
	}

	var user models.User

	if err := databases.DB.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không tìm thấy user id",
		})
		return
	}
	var total_price float64

	var updatedCartItems []models.CartItem
	var cart models.Cart
	if err := databases.DB.Where("id = ?", c.Param("id")).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không tìm thấy cart",
		})
		return
	}

	for _, item := range input.CartItems {
		var product models.Product
		if err := databases.DB.Where("id = ?", item.ProductID).First(&product).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"err": err.Error(),
				"mes": "Lỗi không tìm thấy product id",
			})
			return
		}

		var cartitem models.CartItem

		if err := databases.DB.Where("cart_id = ? and product_id = ?", cart.ID, item.ProductID).First(&cartitem).Error; err != nil {
			cartitem = models.CartItem{
				CartID:    cart.ID,
				ProductID: product.ID,
				Quantity:  item.Quantity,
				Price:     float64(item.Quantity) * product.Price,
			}

			if err := databases.DB.Create(&cartitem).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"err": err.Error(),
					"mes": "Lỗi tạo mới cart item",
				})
				return
			}
		} else {
			cartitem.Quantity = item.Quantity
			cartitem.Price = float64(cartitem.Quantity) * product.Price

			if err := databases.DB.Save(&cartitem).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"err": err.Error(),
					"mes": "Lỗi không lưu được cartitem",
				})
				return
			}
		}

		updatedCartItems = append(updatedCartItems, cartitem)

		total_price += cartitem.Price

	}

	cart.TotalPrice = total_price

	if err := databases.DB.Save(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không lưu được cart",
		})
		return
	}
	//c.JSON(http.StatusOK, cart)
	// Lấy lại cart_items sau khi cập nhật
	var cartItems []models.CartItem
	if err := databases.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không tìm thấy cart items",
		})
		return
	}

	// Tạo cấu trúc phản hồi
	response := gin.H{
		"id":          cart.ID,
		"user_id":     cart.UserID,
		"user":        user,             // Thêm thông tin người dùng
		"cart_items":  updatedCartItems, // Đưa cartItems vào phản hồi
		"total_price": cart.TotalPrice,
		"created_at":  cart.CreatedAt,
	}

	// Phản hồi thành công
	c.JSON(http.StatusOK, response)

}
