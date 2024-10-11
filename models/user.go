package models

import "time"

// người dùng
type User struct {
	ID         int        `json:"id" gorm:"primaryKey"`
	Name       string     `json:"name"`
	Email      string     `json:"email" gorm:"unique"`
	Password   string     `json:"password"`
	Address    string     `json:"address"`
	City       string     `json:"city"`
	PostalCode string     `json:"postal_code"`
	Orders     []Order    `json:"orders" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Một người dùng có nhiều đơn hàng
	CreatedAt  *time.Time `json:"created_at"`
}
