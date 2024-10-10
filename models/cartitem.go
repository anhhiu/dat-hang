package models

type CartItem struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	CartID    uint     `json:"cart_id"`
	Cart      *Cart    `json:"cart" gorm:"foreignKey:CartID"`
	ProductID uint     `json:"product_id"`
	Product   *Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int      `json:"quantity"`
	Price     float64  `json:"price"`
}
