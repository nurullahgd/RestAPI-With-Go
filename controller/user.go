package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurullahgd/RestAPI-With-Go/config"
	"github.com/nurullahgd/RestAPI-With-Go/models"
)

func ListUsers(ctx *gin.Context) {
	users := []models.User{}
	result := config.DB.Find(&users)
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func CreateUser(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)

	if err := config.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusCreated, &user)
	}

}

func UpdateUser(ctx *gin.Context) {
	var user models.User
	userID := ctx.Param("id")
	if err := config.DB.First(&user, userID).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "User not found"})
		return
	}

	ctx.BindJSON(&user)
	if err := config.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, &user)
	}
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	if result := config.DB.First(&user, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	config.DB.Delete(&user)
	ctx.Status(http.StatusOK)

}
