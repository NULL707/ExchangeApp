package controllers

import (
	"ExchangeApp/global"
	"ExchangeApp/models"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var cacheKey = "articles"

func CreateArticle(ctx *gin.Context) {
	var article models.Article
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := global.DB.AutoMigrate(&models.Article{}); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to migrate database"})
		return
	}
	if err := global.DB.Create(&article).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create article"})
		return
	}

	if err := global.RedisDB.Del(cacheKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cache"})
		return
	}

	ctx.JSON(http.StatusCreated, article)
}

func GetArticles(ctx *gin.Context) {

	cachedData, err := global.RedisDB.Get(cacheKey).Result()
	if err == redis.Nil {
		var articles []models.Article
		if err := global.DB.Find(&articles).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "No articles found"})
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get articles"})
				return
			}
		}

		articleJSON, err := json.Marshal(articles)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize articles"})
			return
		}

		if err := global.RedisDB.Set(cacheKey, articleJSON, 10*time.Minute).Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cache articles"})
			return
		}

		ctx.JSON(http.StatusOK, articles)
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get articles from cache"})
		return
	} else {
		var articles []models.Article
		if err := json.Unmarshal([]byte(cachedData), &articles); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deserialize articles"})
			return
		}
		ctx.JSON(http.StatusOK, articles)
	}
}

func GetArticleByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var article models.Article
	if err := global.DB.First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get article"})
		return
	}
	ctx.JSON(http.StatusOK, article)
}
