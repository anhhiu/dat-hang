package models

import "time"

// sản phẩm
type Product struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Stock       int        `json:"stock"`
	CategoryID  uint       `json:"category_id"`
	Category    *Category  `json:"category" gorm:"foreignKey:CategoryID"`
	CreatedAt   *time.Time `json:"created_at"`
}
