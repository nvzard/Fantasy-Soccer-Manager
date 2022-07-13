package service

import (
	"github.com/nvzard/soccer-manager/database"
	"github.com/nvzard/soccer-manager/model"
	"gorm.io/gorm"
)

func GetPlayerByID(playerID string) (model.Player, error) {
	var player model.Player

	if err := database.DB.First(&player, "id = ?", playerID).Error; err != nil {
		return player, err
	}

	return player, nil
}

func UpdatePlayer(player model.Player, playerPatch model.PlayerPatch) (model.Player, error) {
	if playerPatch.FirstName != "" {
		player.FirstName = playerPatch.FirstName
	}
	if playerPatch.LastName != "" {
		player.LastName = playerPatch.LastName
	}
	if playerPatch.Country != "" {
		player.Country = playerPatch.Country
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Save(&player)
		if err := result.Error; err != nil {
			logger.Errorw("Failed to update player", "error", err, "player", player)
			return err
		}
		return nil
	})

	return player, err
}
