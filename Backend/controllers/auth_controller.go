package controllers

import (
	"ExchangeApp/global"
	"ExchangeApp/models"
	"ExchangeApp/utils"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = hashedPassword

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	if err := global.DB.AutoMigrate(&models.User{}); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to migrate database"})
		return
	}
	if err := global.DB.Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(200, gin.H{"token": token})
}

func Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := global.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}
	if !utils.CheckPassword(input.Password, user.Password) {
		ctx.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}
	ctx.JSON(200, gin.H{"token": token})
}
