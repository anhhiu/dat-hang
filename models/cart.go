package models

import "time"

//giỏ hàng
type Cart struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	UserID     uint       `json:"user_id"`
	User       *User      `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CartItems  []CartItem `json:"cart_items" gorm:"foreignKey:CartID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TotalPrice float64    `json:"total_price"`
	CreatedAt  *time.Time `json:"created_at"`
}
