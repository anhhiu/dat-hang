package models

import "time"

//phương thức giao hàng/ vận chuyển
type ShippingMethod struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Orders      []Order    `json:"orders" gorm:"foreignKey:ShippingMethodID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   *time.Time `json:"created_at"`
}
