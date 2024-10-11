package controllers

import (
	"dathang/databases"
	"dathang/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create product
// @Tags Product
// @Param product body models.Product true "Product data"
// @Router /product/ [post]
func CreateProduct(c *gin.Context) {
	var input struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Hinh        string  `json:"hinh"`
		Stock       int     `json:"stock"`
		CategoryID  uint    `json:"category_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	createat := time.Now()
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		CategoryID:  input.CategoryID,
		CreatedAt:   &createat,
	}

	if err := databases.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, product)
}

// @Summary Get all product
// @Tags Product
// @Router /product/ [get]
func GetAllProduct(c *gin.Context) {
	var products []models.Product

	if err := databases.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

// @Summary Get product by id
// @Tags Product
// @Param id path int true "ID"
// @Router /product/{id} [get]
func GetProductById(c *gin.Context) {
	var product models.Product

	if err := databases.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)

}

// @Summary Update product by id
// @Tags Product
// @Param id path int true "ID"
// @Param product body models.Product true "Product data"
// @Router /product/{id} [put]
func UpdateProduct(c *gin.Context) {
	var input struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Hinh        string  `json:"hinh"`
		Stock       int     `json:"stock"`
		CategoryID  uint    `json:"category_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product

	if err := databases.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	updateat := time.Now()

	product.Name = input.Name
	product.Description = input.Description
	product.Hinh = input.Hinh
	product.Price = input.Price
	product.Stock = input.Stock
	product.CategoryID = input.CategoryID
	product.UpdatedAt = &updateat

	if err := databases.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary Delete product by id
// @Tags Product
// @Param id path int true "ID"
// @Router /product/{id} [delete]
func DeleteProductById(c *gin.Context) {
	var product models.Product

	if err := databases.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := databases.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Xóa thành công.")

}
