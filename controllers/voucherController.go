package controllers

import (
	"dathang/databases"
	"dathang/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create voucher
// @Tags Voucher
// @Param voucher body models.Voucher true "Voucher data"
// @Router /voucher/ [post]
func CreateVoucher(c *gin.Context) {
	var input struct {
		Code     string  `json:"code"`
		Discount float64 `json:"discount"`
		Expiry   string  `json:"expiry"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"mes": "Lỗi đầu vào json",
		})
		return
	}

	expiry, err := time.Parse("02/01/2006", input.Expiry)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"mes": "Lỗi định dạng ngày",
		})
		return
	}

	createdAt := time.Now()

	voucher := models.Voucher{
		Code:      input.Code,
		Discount:  input.Discount,
		Expiry:    &expiry,
		CreatedAt: &createdAt,
	}

	if err := databases.DB.Create(&voucher).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không tạo được voucher",
		})
		return
	}

	c.JSON(http.StatusCreated, voucher)

}

// @Summary Get ALL voucher
// @Tags Voucher
// @Router /voucher/ [get]
func GetAllVoucher(c *gin.Context) {
	var vouchers []models.Voucher
	if err := databases.DB.Find(&vouchers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi truy vấn voucher",
		})
		return
	}

	c.JSON(http.StatusOK, vouchers)
}

// @Summary Get voucher by id
// @Tags Voucher
// @Param id path int true "ID"
// @Router /voucher/{id} [get]
func GetVoucherById(c *gin.Context) {
	var voucher models.Voucher

	if err := databases.DB.Where("id = ?", c.Param("id")).First(&voucher).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không tìm thấy",
		})
		return
	}

	c.JSON(http.StatusOK, voucher)
}

// @Summary Delete voucher
// @Tags Voucher
// @Param id path int true "ID"
// @Router /voucher/{id} [delete]
func DeleteVoucher(c *gin.Context) {
	var voucher models.Voucher
	if err := databases.DB.Where("id = ?", c.Param("id")).First(&voucher).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không tìm thấy id",
		})
		return
	}
	if err := databases.DB.Delete(&voucher).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không xóa được trong database",
		})
		return
	}

	c.JSON(http.StatusOK, "xóa thành công")
}

// @Summary Update voucher
// @Tags Voucher
// @Param id path int true "ID"
// @Param voucher body models.Voucher true "Voucher data"
// @Router /voucher/{id} [put]
func UpdateVoucher(c *gin.Context) {
	var input struct {
		Code     string  `json:"code"`
		Discount float64 `json:"discount"`
		Expiry   string  `json:"expiry"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"mes": "Lỗi đầu vào json",
		})
		return
	}

	expiry, err := time.Parse("02/01/2006", input.Expiry)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"mes": "Lỗi định dạng ngày",
		})
		return
	}

	var voucher models.Voucher

	if err := databases.DB.Where("id = ?", c.Param("id")).First(&voucher).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không tìm thấy",
		})
		return
	}

	voucher.Code = input.Code
	voucher.Discount = input.Discount
	voucher.Expiry = &expiry

	if err := databases.DB.Save(&voucher).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"mes": "Lỗi không lưu được vào database",
		})
		return
	}

	c.JSON(http.StatusOK, voucher)
}
