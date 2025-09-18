package controllers

import (
	"ExchangeApp/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func LikeArticle(ctx *gin.Context) {
	// Implementation for liking an article
	articleId := ctx.Param("id")

	likeKey := "article:" + articleId + ":likes"

	err := global.RedisDB.Incr(likeKey).Err()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Article liked"})
}

func GetArticleLikes(ctx *gin.Context) {
	articleId := ctx.Param("id")
	likeKey := "article:" + articleId + ":likes"
	likes, err := global.RedisDB.Get(likeKey).Result()
	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
