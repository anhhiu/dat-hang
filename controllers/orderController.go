package controllers

import (
	"dathang/databases"
	"dathang/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type InputOrder struct {
	UserID           uint                   `json:"user_id"`
	CartID           uint                   `json:"cart_id"`
	StatusID         uint                   `json:"status_id"`
	ShippingAddress  models.ShippingAddress `json:"shipping_address"`
	ShippingMethodID uint                   `json:"shipping_method_id"`
	VoucherID        *uint                  `json:"voucher_id"`
	PaymentStatus    string                 `json:"payment_status"` // pending, completed, failed
}

/* func CalculateTotalPrice(cart models.Cart, voucher *models.Voucher) float64 {
	totalPrice := 0.0

	for _, item := range cart.CartItems {
		totalPrice += item.Price * float64(item.Quantity)
		fmt.Printf("Sản phẩm: %s, Giá: %f, Số lượng: %d\n", item.Product.Name, item.Product.Price, item.Quantity)
	}

	if voucher != nil {
		totalPrice -= voucher.Discount
	}
	fmt.Printf("sum : %.2f ", totalPrice)
	return totalPrice
} */

func CalculateTotalPrice(cart models.Cart, voucher *models.Voucher) float64 {
	totalPrice := 0.0

	// Tính tổng giá trị của các sản phẩm trong giỏ hàng
	for _, item := range cart.CartItems {
		price := item.Product.Price // Lấy giá từ sản phẩm
		totalPrice += price * float64(item.Quantity)
		fmt.Printf("Sản phẩm: %s, Giá: %.2f, Số lượng: %d\n", item.Product.Name, price, item.Quantity)
	}

	// In ra tổng trước khi áp dụng voucher
	fmt.Printf("Tổng giá trị trước khi áp dụng voucher: %.2f\n", totalPrice)

	// Áp dụng voucher nếu có
	if voucher != nil {
		totalPrice -= voucher.Discount
	}

	// Đảm bảo tổng tiền không âm
	if totalPrice < 0 {
		totalPrice = 0
	}

	fmt.Printf("Tổng sau khi áp dụng voucher: %.2f\n", totalPrice)
	return totalPrice
}

// @Summary Create Order
// @Tags Order
// @Param order body models.Order true "Order data"
// @Router /order/ [post]
func CreateOrder(c *gin.Context) {
	var input InputOrder

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"mes": "lỗi định dạng json",
		})
		return
	}

	var cart models.Cart

	if err := databases.DB.Preload("CartItems.Product").First(&cart, input.CartID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không tìm thấy  cart",
		})
		return
	}

	var voucher *models.Voucher
	if input.VoucherID != nil {
		voucher := &models.Voucher{}

		if err := databases.DB.First(voucher, *input.VoucherID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"err": err.Error(),
				"mes": "voucher not found",
			})
			return
		}
	}
	toTalPrice := CalculateTotalPrice(cart, voucher)
	createdAt := time.Now()
	order := models.Order{
		UserID:           input.UserID,
		CartID:           input.CartID,
		VoucherID:        *input.VoucherID,
		StatusID:         input.StatusID,
		ShippingAddress:  &input.ShippingAddress,
		ShippingMethodID: input.ShippingMethodID,
		TotalPrice:       toTalPrice,
		PaymentStatus:    input.PaymentStatus,
		CreatedAt:        &createdAt,
	}

	if err := databases.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không tạo được đơn hàng",
		})
		return
	}

	for _, cartItem := range cart.CartItems {
		product := cartItem.Product

		if product.Stock < cartItem.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "không đủ hàng trong kho " + product.Name,
			})
			return
		}
		product.Stock -= cartItem.Quantity

		if err := databases.DB.Save(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
				"mes": "Lỗi không lưu được các sản phẩm vừa cập nhật",
			})
			return
		}
	}
	// xóa voucher
	/* if voucher != nil {
		if err := databases.DB.Delete(&voucher).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
				"mes": "Lỗi khi xóa voucher",
			})
			return
		}
	} */
	//xóa các bảng liên quan
	/*
	   	if err := databases.DB.Where("cart_id = ?", cart.ID).Delete(&models.Order{}).Error; err != nil {
	   		c.JSON(http.StatusInternalServerError, gin.H{
	   			"err": err.Error(),
	   			"mes": " Lỗi khi xóa các bảng liên quan",
	   		})
	   		return
	   	}
	   // xóa giỏ hàng khi đã đặt đơn
	   	if err := databases.DB.Delete(&cart).Error; err != nil {
	   		c.JSON(http.StatusInternalServerError, gin.H{
	   			"err": err.Error(),
	   			"mes": "Lỗi khi xóa cart",
	   		})
	   		return
	   	} */

	fmt.Printf("Đơn hàng vừa tạo: %+v\n", order)
	c.JSON(http.StatusOK, gin.H{
		"mes":  "Tạo đơn hàng thành công",
		"data": order,
	})

}

// @Summary Get All Order
// @Tags Order
// @Router /order/ [get]
func GetAllOrder(c *gin.Context) {
	var orders []models.Order

	if err := databases.DB.Preload("Cart.CartItems.Product").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không thể truy xuất vào db",
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}
