package models

import "time"

// mã khuyến mại
type Voucher struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Code      string     `json:"code"`
	Discount  float64    `json:"discount"`
	Expiry    *time.Time `json:"expiry"`
	Orders    []Order    `json:"orders" gorm:"foreignKey:VoucherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt *time.Time `json:"created_at"`
}
