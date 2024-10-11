package controllers

import (
	"dathang/databases"
	"dathang/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create users
// @Tags User
// @Param user body models.User true "User data"
// @Router /user/ [post]
func CreateUser(c *gin.Context) {
	var input struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		Address    string `json:"address"`
		City       string `json:"city"`
		PostalCode string `json:"postal_code"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	createat := time.Now()
	user := models.User{
		Name:       input.Name,
		Email:      input.Email,
		Password:   input.Password,
		Address:    input.Address,
		City:       input.City,
		PostalCode: input.PostalCode,
		CreatedAt:  &createat,
	}

	if err := databases.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}

// @Summary Get all users
// @Tags User
// @Router /user/ [get]
func GetAllUser(c *gin.Context) {
	var users []models.User

	if err := databases.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

// @Summary Get users by id
// @Tags User
// @Param id path int true "ID"
// @Router /user/{id} [get]
func GetUserById(c *gin.Context) {
	var user models.User

	if err := databases.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
