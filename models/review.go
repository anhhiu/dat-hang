package models

import "time"

// đánh giá // chua su dung
type Review struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	ProductID int        `json:"product_id"`
	UserID    int        `json:"user_id"`
	Rating    int        `json:"rating"` // 1-5 sao
	Comment   string     `json:"comment"`
	Product   *Product   `json:"product" gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE;"`
	User      *User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	CreatedAt *time.Time `json:"created_at"`
}
