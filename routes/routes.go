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

	review := r.Group("/api/review")
	{
		review.GET("/", controllers.GetAllReview)
		review.POST("/", controllers.CreateReview)
		review.GET("/:id", controllers.GetReviewById)
		review.PUT("/:id", controllers.UpdateReview)
		review.DELETE("/:id", controllers.DeleteReview)
	}

	cartitem := r.Group("/api/cartitem")
	{
		cartitem.GET("/", controllers.GetAllCartItem)
		cartitem.POST("/", controllers.CreateCartItem)
		//cartitem.GET("/:id", controllers.GetReviewById)
		//cartitem.PUT("/:id", controllers.UpdateReview)
		//cartitem.DELETE("/:id", controllers.DeleteReview)
	}

	cart := r.Group("/api/cart")
	{
		cart.GET("/", controllers.GetAllCart)
		cart.POST("/", controllers.CreateCart)
		//cart.GET("/:id", controllers.GetReviewById)
		//cart.PUT("/:id", controllers.UpdateReview)
		//cart.DELETE("/:id", controllers.DeleteReview)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
