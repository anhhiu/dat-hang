package models

import "time"

// lịch sử đặt hàng
type OrderHistory struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	OrderID   uint       `json:"order_id"`
	Order     *Order     `json:"order" gorm:"foreignKey:OrderID"`
	Status    string     `json:"status"` // pending, shipped, delivered, cancelled
	UpdatedAt *time.Time `json:"updated_at"`
}
