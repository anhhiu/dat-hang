package controllers

import (
	"dathang/databases"
	"dathang/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create Categorys
// @Tags Category
// @Param category body models.Category true "Category data"
// @Router /category/ [post]
func CreateCategory(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Products []struct {
			Name        string  `json:"name"`
			Description string  `json:"description"`
			Price       float64 `json:"price"`
			Stock       int     `json:"stock"`
			Hinh        string  `json:"hinh"`
		} `json:"products"`
	}

	// Xử lý dữ liệu đầu vào
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Tạo Category mới
	now := time.Now()
	category := models.Category{
		Name:      input.Name,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	// Lưu Category trước
	if err := databases.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Lưu các Products kèm theo CategoryID
	for _, p := range input.Products {
		product := models.Product{
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			Hinh:        p.Hinh,
			CategoryID:  category.ID, // Gán CategoryID
			CreatedAt:   &now,
			UpdatedAt:   &now,
		}

		// Lưu từng Product
		if err := databases.DB.Create(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	// Tải lại Category với các Products
	if err := databases.DB.Preload("Products").First(&category, category.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Trả về Category với các sản phẩm đã lưu
	c.JSON(http.StatusCreated, category)
}

// @Summary Get all Categorys
// @Tags Category
// @Router /category/ [get]
func GetAllCategory(c *gin.Context) {
	var categorys []models.Category

	if err := databases.DB.Preload("Products").Find(&categorys).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, categorys)
}

// @Summary Get Categorys by id
// @Tags Category
// @Param id path int true "ID"
// @Router /category/{id} [get]
func GetCategoryById(c *gin.Context) {
	var category models.Category

	if err := databases.DB.Preload("Products").Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, category)
}

// @Summary Update Categorys by id
// @Tags Category
// @Param id path int true "ID"
// @Param category body models.Category true "Category data"
// @Router /category/{id} [put]
func UpdatedCategory(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Products []struct {
			Name        string  `json:"name"`
			Hinh        string  `json:"hinh"`
			Price       float64 `json:"price"`
			Stock       int     `json:"stock"`
			Description string  `json:"description"`
		}
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var category models.Category

	if err := databases.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	updataat := time.Now()

	category.Name = input.Name
	category.UpdatedAt = &updataat

	if err := databases.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	for _, p := range input.Products {

		var product models.Product

		if err := databases.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		product.Name = p.Name
		product.Description = p.Description
		product.Hinh = p.Hinh
		product.CategoryID = category.ID
		product.Price = p.Price
		product.Stock = p.Stock
		product.UpdatedAt = &updataat

		if err := databases.DB.Save(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		//c.JSON(http.StatusOK, product)
	}

	c.JSON(http.StatusOK, category)
}

// @Summary Delete Categorys by id
// @Tags Category
// @Param id path int true "ID"
// @Router /category/{id} [delete]
func DeleteCategoryById(c *gin.Context) {
	var category models.Category

	if err := databases.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := databases.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Đã xóa thành công.")
}
