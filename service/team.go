package service

import (
	"github.com/nvzard/soccer-manager/database"
	"github.com/nvzard/soccer-manager/model"
	"gorm.io/gorm"
)

func GetTeamByID(teamID string) (model.Team, error) {
	var team model.Team

	if err := database.DB.Preload("Players").First(&team, "id = ?", teamID).Error; err != nil {
		return team, err
	}

	return team, nil
}

func UpdateTeam(team model.Team, teamPatch model.TeamPatch) (model.Team, error) {
	if teamPatch.Name != "" {
		team.Name = teamPatch.Name
	}
	if teamPatch.Country != "" {
		team.Country = teamPatch.Country
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Save(&team)
		if err := result.Error; err != nil {
			logger.Errorw("Failed to update team", "error", err, "team", team)
			return err
		}
		return nil
	})

	return team, err
}
