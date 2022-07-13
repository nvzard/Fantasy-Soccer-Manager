package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvzard/soccer-manager/service"
)

func GetTeam(context *gin.Context) {
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
		"team_value":     team.CalculateTeamValue(),
		"owner_id":       team.UserID,
		"players":        team.Players,
	})
}
