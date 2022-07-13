package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nvzard/soccer-manager/auth"
	"github.com/nvzard/soccer-manager/model"
	"github.com/nvzard/soccer-manager/service"
)

func CreateTransfer(context *gin.Context) {
	var transferRequest model.TransferRequest
	if err := context.ShouldBindJSON(&transferRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := transferRequest.Validate(); err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// fmt.Printf("%+v \n", transferRequest)
	playerID := strconv.FormatUint(uint64(transferRequest.PlayerID), 10)
	player, err := service.GetPlayerByID(playerID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "player does not exist"})
		context.Abort()
		return
	}

	// auth for player
	userAuth, exists := auth.GetUserAuth(context)
	if !exists || userAuth.TeamID != player.TeamID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to transfer this player"})
		context.Abort()
		return
	}

	// check if player is already on transfer list
	transfer, _ := service.GetTransferByPlayerID(playerID)
	if transfer.ID != 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "player already on transfer list"})
		context.Abort()
		return
	}

	transfer, err = service.CreateTransfer(transferRequest, player)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"transfer": transfer})
}

// func GetUser(context *gin.Context) {
// 	userEmail := context.Params.ByName("email")

// 	user, err := service.GetUser(userEmail)

// 	if err != nil {
// 		context.JSON(http.StatusNotFound, gin.H{"error": "user does not exist"})
// 		context.Abort()
// 		return
// 	}

// 	context.JSON(http.StatusCreated, gin.H{
// 		"id":         user.ID,
// 		"email":      user.Email,
// 		"first_name": user.FirstName,
// 		"last_name":  user.LastName,
// 		"team_id":    user.TeamID,
// 	})
// }
