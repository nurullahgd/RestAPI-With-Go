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

}
