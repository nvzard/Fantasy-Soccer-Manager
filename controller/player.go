package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvzard/soccer-manager/auth"
	"github.com/nvzard/soccer-manager/model"
	"github.com/nvzard/soccer-manager/service"
)

func GetPlayer(context *gin.Context) {
	playerID := context.Params.ByName("id")
	player, err := service.GetPlayerByID(playerID)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "player does not exist"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":           player.ID,
		"first_name":   player.FirstName,
		"last_name":    player.LastName,
		"country":      player.Country,
		"age":          player.Age,
		"position":     player.Position,
		"market_value": player.MarketValue,
		"team_id":      player.TeamID,
	})
}

func UpdatePlayer(context *gin.Context) {
	playerID := context.Params.ByName("id")
	player, err := service.GetPlayerByID(playerID)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "player does not exist"})
		context.Abort()
		return
	}

	userAuth, exists := auth.GetUserAuth(context)
	if !exists || userAuth.TeamID != player.TeamID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to update this player"})
		context.Abort()
		return
	}

	var playerPatch model.PlayerPatch
	if err := context.ShouldBindJSON(&playerPatch); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	player, err = service.UpdatePlayer(player, playerPatch)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update player"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":           player.ID,
		"first_name":   player.FirstName,
		"last_name":    player.LastName,
		"country":      player.Country,
		"age":          player.Age,
		"position":     player.Position,
		"market_value": player.MarketValue,
		"team_id":      player.TeamID,
	})
}
