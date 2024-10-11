package models

type CartItem struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	CartID uint `json:"cart_id"`
	//Cart      *Cart    `json:"cart" gorm:"foreignKey:CartID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ProductID uint     `json:"product_id"`
	Product   *Product `json:"product" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Quantity  int      `json:"quantity"`
	Price     float64  `json:"price"`
}
