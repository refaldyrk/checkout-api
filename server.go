package main

import (
	"checkout-api/config"
	"checkout-api/handler"
	"checkout-api/middleware"
	"checkout-api/repository"
	"checkout-api/service"
	"checkout-api/utils/data"
	"checkout-api/utils/http_response"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupServer(router *gin.Engine) {
	responseWriter := http_response.NewResponseWriter()
	database, err := config.NewDb(config.Envs.DB_URI)
	if err != nil {
		panic(err)
	}

	// repositories
	userRepo := repository.NewUserRepo(database)
	productRepo := repository.NewProductRepo(database)
	cartRepo := repository.NewCartRepo(database)
	productOrderRepo := repository.NewProductOrderRepo(database)

	// seed data
	data.SeedData(productRepo)
	data.SeedSuperuser(userRepo)

	// services
	authService := service.NewAuthService(userRepo)
	productService := service.NewProductService(productRepo, cartRepo, productOrderRepo)
	cartService := service.NewCartService(cartRepo, productRepo, productOrderRepo)

	// handlers
	authHandler := handler.NewAuthHandler(responseWriter, authService)
	productHandler := handler.NewProductHandler(responseWriter, productService)
	cartHandler := handler.NewCartHandler(responseWriter, cartService)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error":   false,
			"message": "pong!",
		})
	})
	router.POST("/auth/login", authHandler.Login)
	router.POST("/auth/check-token", authHandler.CheckToken)
	router.POST("/auth/register", authHandler.Register)

	secureRouter := router.Group("/")
	{
		secureRouter.Use(middleware.JWTMiddleware(responseWriter, authService))

		secureRouter.GET("/products", productHandler.GetList)
		secureRouter.GET("/products/category-list", productHandler.GetCategoryList)
		secureRouter.POST("/products/:id/add-to-cart", productHandler.AddItemToCart)
		secureRouter.GET("/cart/orders", cartHandler.GetCartItems)
		secureRouter.DELETE("/cart/orders/:id", cartHandler.RemoveItemFromCart)
		secureRouter.POST("/cart/checkout", cartHandler.CheckoutCart)
	}

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
