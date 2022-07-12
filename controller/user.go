package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvzard/soccer-manager/model"
	"github.com/nvzard/soccer-manager/service"
)

func RegisterUser(context *gin.Context) {
	var newUser model.User
	if err := context.ShouldBindJSON(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := newUser.HashPassword(newUser.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	newUser, err := service.CreateUser(newUser)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userID": newUser.ID, "email": newUser.Email})
}
