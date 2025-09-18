package router

import (
	"ExchangeApp/controllers"
	"ExchangeApp/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
	api := r.Group("/api")
	api.GET("/exchange_rates", controllers.GetExchangeRate)
	api.Use(middlewares.AuthMiddleware())
	{
		api.POST("/exchange_rates", controllers.CreateExchangeRate)
		api.POST("/article", controllers.CreateArticle)
		api.GET("/article/:id", controllers.GetArticleByID)
		api.GET("/articles", controllers.GetArticles)

		api.POST("/article/:id/like", controllers.LikeArticle)
		api.GET("/article/:id/likes", controllers.GetArticleLikes)
	}
	return r
}
