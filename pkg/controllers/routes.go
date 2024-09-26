package controllers

import (
	"products/configs"
	_ "products/docs"
	 "github.com/swaggo/gin-swagger"
     "github.com/swaggo/files" 
	"github.com/gin-gonic/gin"
)

func RunRoutes() *gin.Engine {
	router := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", Ping)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}
	apiG := router.Group("/api", checkUserAuthentication)
	userG := apiG.Group("/users")
	{
		userG.GET("", GetAllUsers)
		userG.GET("/:id", GetUserByID)
		userG.POST("", CreateUser)
		userG.PUT("/:id", UpdateUserByID)
	}

	productG := apiG.Group("/product")
	{
		productG.GET("", GetAllProducts)
		productG.GET("/:id", GetProductByID)
		productG.POST("", CreateProduct)
		productG.DELETE("/:id", DeleteProductByID)
	}

	categoryG := apiG.Group("/category")
	{
		categoryG.GET("", GetAllCategories)
		categoryG.GET("/:id", GetCategoryByID)
		categoryG.POST("", CreateCategory)
		categoryG.PUT("/:id", UpdateCategoryByID)
		categoryG.DELETE("/:id", DeleteCategoryByID)
	}

	wishlistG := apiG.Group("/wishlist")
	{
		wishlistG.GET("/:id", GetWishlistByUserID)
		wishlistG.POST("", AddToWishlist)
		wishlistG.DELETE("/:id", RemoveFromWishlist)
	}

	orderG := apiG.Group("/order")
	{
		orderG.GET("/:id", GetUserOrderByID)
		orderG.POST("", CreateOrder)
		orderG.PATCH("/:id", RemoveItem)
		orderG.GET("", GetAllOrders)
	}
	
	return router
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
