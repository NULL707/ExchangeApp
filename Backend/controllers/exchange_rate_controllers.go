package controllers

import (
	"ExchangeApp/global"
	"ExchangeApp/models"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateExchangeRate(ctx *gin.Context) {
	var exchangeRate models.ExchangeRate
	if err := ctx.ShouldBindJSON(&exchangeRate); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	exchangeRate.Date = time.Now()

	if err := global.DB.AutoMigrate(&models.ExchangeRate{}); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to migrate database"})
		return
	}
	if err := global.DB.Create(&exchangeRate).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create exchange rate"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Exchange rate created successfully"})
}

func GetExchangeRate(ctx *gin.Context) {
	var exchangeRates []models.ExchangeRate
	if err := global.DB.Find(&exchangeRates).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(404, gin.H{"error": "No exchange rates found"})
			return
		}
		ctx.JSON(500, gin.H{"error": "Failed to fetch exchange rates"})
		return
	}
	ctx.JSON(200, gin.H{"exchange_rates": exchangeRates})
}
