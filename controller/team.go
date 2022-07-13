package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvzard/soccer-manager/model"
	"github.com/nvzard/soccer-manager/service"
)

func GetTeam(context *gin.Context) {
	teamID := context.Params.ByName("id")
	team, err := service.GetTeamByID(teamID)

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

func UpdateTeam(context *gin.Context) {
	teamID := context.Params.ByName("id")
	team, err := service.GetTeamByID(teamID)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "team does not exist"})
		context.Abort()
		return
	}

	var teamPatch model.TeamPatch
	if err := context.ShouldBindJSON(&teamPatch); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	team, err = service.UpdateTeam(team, teamPatch)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update player"})
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
