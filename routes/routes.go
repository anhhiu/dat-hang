package routes

import (
	"dathang/controllers"
	"dathang/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	user := r.Group("/api/user")
	{
		user.GET("/", controllers.GetAllUser)
		user.POST("/", controllers.CreateUser)
		user.GET("/:id", controllers.GetUserById)
	}

	category := r.Group("/api/category")
	{
		category.GET("/", controllers.GetAllCategory)
		category.POST("/", controllers.CreateCategory)
		category.GET("/:id", controllers.GetCategoryById)
		category.PUT("/:id", controllers.UpdatedCategory)
		category.DELETE("/:id", controllers.DeleteCategoryById)

	}

	product := r.Group("/api/product")
	{
		product.GET("/", controllers.GetAllProduct)
		product.POST("/", controllers.CreateProduct)
		product.GET("/:id", controllers.GetProductById)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProductById)

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
