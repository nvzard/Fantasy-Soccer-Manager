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

	if err := newUser.Validate(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	context.JSON(http.StatusCreated, gin.H{"id": newUser.ID, "email": newUser.Email})
}

func GetUser(context *gin.Context) {
	userEmail := context.Params.ByName("email")

	user, err := service.GetUser(userEmail)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "user does not exist"})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"team_id":    user.TeamID,
	})
}
