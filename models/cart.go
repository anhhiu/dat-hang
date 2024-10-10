package models

import "time"

//giỏ hàng - 1
type Cart struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	UserID     uint       `json:"user_id"`
	User       *User      `json:"user" gorm:"foreignKey:UserID"`
	CartItems  []CartItem `json:"cart_items" gorm:"foreignKey:CartID"`
	TotalPrice float64    `json:"total_price"`
	CreatedAt  *time.Time `json:"created_at"`
}


