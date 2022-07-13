package service

import (
	"github.com/nvzard/soccer-manager/database"
	"github.com/nvzard/soccer-manager/model"
)

func GetPlayerByID(playerID string) (model.Player, error) {
	var player model.Player

	if err := database.DB.First(&player, "id = ?", playerID).Error; err != nil {
		return player, err
	}

	return player, nil
}
