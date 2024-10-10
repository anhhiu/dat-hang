package models

//chi tiết đặt hàng
type OrderDetail struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	OrderID   uint     `json:"order_id"`
	Order     *Order   `json:"order" gorm:"foreignKey:OrderID"`
	ProductID uint     `json:"product_id"`
	Product   *Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int      `json:"quantity"`
	Price     float64  `json:"price"`
}
