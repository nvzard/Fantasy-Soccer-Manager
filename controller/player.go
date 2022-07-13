package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		"name":         player.FirstName,
		"country":      player.Country,
		"market_value": player.MarketValue,
		"team_id":      player.TeamID,
	})
}

func GetPlayersByTeamID(context *gin.Context) {
	teamID := context.Params.ByName("id")
	team, err := service.GetTeam(teamID)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "team does not exist"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":             team.ID,
		"name":           team.Name,
		"country":        team.Country,
		"available_cash": team.AvailableCash,
		"owner_id":       team.UserID,
	})
}
