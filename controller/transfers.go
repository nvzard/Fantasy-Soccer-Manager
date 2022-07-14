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

	// authentication of request
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
	context.JSON(http.StatusCreated, gin.H{"transfer_id": transfer.ID})
}

func BuyPlayer(context *gin.Context) {
	playerID := context.Params.ByName("player_id")
	player, err := service.GetPlayerByID(playerID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "player does not exist"})
		context.Abort()
		return
	}

	// authentication of request
	userAuth, exists := auth.GetUserAuth(context)
	if !exists || userAuth.TeamID == player.TeamID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "player is already on your team"})
		context.Abort()
		return
	}

	// check if player is on transfer list
	transfer, _ := service.GetTransferByPlayerID(playerID)
	if transfer.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "player not available for transfer"})
		context.Abort()
		return
	}

	// check if user has enough money
	buyersTeamID := strconv.FormatUint(uint64(userAuth.TeamID), 10)
	team, err := service.GetTeamByID(buyersTeamID)
	if err != nil || team.AvailableCash < transfer.AskedPrice {
		context.JSON(http.StatusBadRequest, gin.H{"error": "not enough money"})
		context.Abort()
		return
	}

	// transfer player: close the deal
	err = service.TransferPlayer(player, transfer, userAuth.TeamID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "done"})
}

func GetTransfers(context *gin.Context) {
	transfers, err := service.GetAllPendingTransfers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"transfer_list": transfers,
	})
}
