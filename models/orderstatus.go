package models

import "time"

// trạng thái đơn hàng
type OrderStatus struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Status    string     `json:"status"` // pending, shipped, delivered, cancelled
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
