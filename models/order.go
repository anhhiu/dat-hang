package models

import "time"

//đặt hàng
type Order struct {
	ID               uint             `json:"id" gorm:"primaryKey"`
	UserID           uint             `json:"user_id"`
	User             User             `json:"user" gorm:"foreignKey:UserID"`
	CartID           uint             `json:"cart_id"`
	Cart             *Cart            `json:"cart" gorm:"foreignKey:CartID"`
	StatusID         uint             `json:"status_id"`
	Status           *OrderStatus     `json:"status" gorm:"foreignKey:StatusID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ShippingAddress  *ShippingAddress `json:"shipping_address" gorm:"embedded"`
	ShippingMethodID uint             `json:"shipping_method_id"`
	ShippingMethod   *ShippingMethod  `json:"shipping_method" gorm:"foreignKey:ShippingMethodID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	VoucherID        uint             `json:"voucher_id"`
	Voucher          *Voucher         `json:"voucher" gorm:"foreignKey:VoucherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TotalPrice       float64          `json:"total_price"`
	PaymentStatus    string           `json:"payment_status"` // pending, completed, failed
	OrderDetails     []OrderDetail    `json:"order_details" gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt        *time.Time       `json:"created_at"`
}
