package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvzard/soccer-manager/model"
	"github.com/nvzard/soccer-manager/service"
)

func RegisterUser(context *gin.Context) {
	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	user, err := service.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userID": user.ID, "email": user.Email})
}
