package models

import "time"

// thanh toán
type Payment struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	OrderID   uint       `json:"order_id"`
	Order     *Order     `json:"order" gorm:"foreignKey:OrderID"`
	Amount    float64    `json:"amount"`
	Method    string     `json:"method"` // Ví điện tử, thẻ, tiền mặt
	Status    string     `json:"status"` // pending, completed, failed
	CreatedAt *time.Time `json:"created_at"`
}
