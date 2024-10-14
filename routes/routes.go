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
		cartitem.GET("/:id", controllers.GetCartItemById)
		cartitem.PUT("/:id", controllers.UpdateCartItem)
		cartitem.DELETE("/:id", controllers.DeleteCartItem)
	}

	cart := r.Group("/api/cart")
	{
		cart.GET("/", controllers.GetAllCart)
		cart.POST("/", controllers.CreateCart)
		cart.GET("/:id", controllers.GetCartById)
		cart.PUT("/:id", controllers.UpdateCart)
		cart.DELETE("/:id", controllers.DeleteCart)
	}

	voucher := r.Group("/api/voucher")
	{
		voucher.GET("/", controllers.GetAllVoucher)
		voucher.POST("/", controllers.CreateVoucher)
		voucher.GET("/:id", controllers.GetVoucherById)
		voucher.PUT("/:id", controllers.UpdateVoucher)
		voucher.DELETE("/:id", controllers.DeleteVoucher)
	}

	shippingmethod := r.Group("/api/shippingmethod")
	{
		shippingmethod.GET("/", controllers.GetAllShippingMethod)
		shippingmethod.POST("/", controllers.CreateShippingMethod)
	}
	orderstatus := r.Group("/api/orderstatus")
	{
		orderstatus.GET("/", controllers.GetAllOrderStatus)
		orderstatus.POST("/", controllers.CreateOrderStatus)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
