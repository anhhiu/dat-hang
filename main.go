package main

import (
	"dathang/databases"
	"dathang/models"
	"dathang/routes"
	"log"
)

func main() {
	databases.ConnectDatabase()
	databases.DB.AutoMigrate(&models.User{})
	databases.DB.AutoMigrate(&models.Category{})
	databases.DB.AutoMigrate(&models.Product{})
	databases.DB.AutoMigrate(&models.Cart{})
	databases.DB.AutoMigrate(&models.CartItem{})
	databases.DB.AutoMigrate(&models.ShippingAddress{})
	databases.DB.AutoMigrate(&models.OrderStatus{})
	databases.DB.AutoMigrate(&models.ShippingMethod{})
	databases.DB.AutoMigrate(&models.Order{})
	databases.DB.AutoMigrate(&models.OrderDetail{})
	databases.DB.AutoMigrate(&models.Payment{})
	databases.DB.AutoMigrate(&models.Review{})
	databases.DB.AutoMigrate(&models.Voucher{})
	databases.DB.AutoMigrate(&models.OrderHistory{})

	r := routes.SetUpRouter()

	if err := r.Run(":9999"); err != nil {
		log.Fatalf("loi roi %v", err.Error())
	}
}
