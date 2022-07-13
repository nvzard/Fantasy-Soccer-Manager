package service

import (
	"github.com/nvzard/soccer-manager/database"
	"github.com/nvzard/soccer-manager/model"
)

func GetTeam(teamID string) (model.Team, error) {
	var team model.Team

	if err := database.DB.Preload("Players").First(&team, "id = ?", teamID).Error; err != nil {
		return team, err
	}

	return team, nil
}
