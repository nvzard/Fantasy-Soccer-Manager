package service

import (
	"github.com/nvzard/soccer-manager/database"
	"github.com/nvzard/soccer-manager/model"
	"gorm.io/gorm"
)

func CreateTransfer(transferRequest model.TransferRequest, player model.Player) (model.Transfer, error) {
	newTransfer := model.Transfer{
		PlayerID:    player.ID,
		AskedPrice:  transferRequest.AskedPrice,
		MarketValue: player.MarketValue,
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&newTransfer)
		if err := result.Error; err != nil {
			logger.Errorw("Failed to transfer player", "error", err, "transfer", newTransfer)
			return err
		}
		return nil
	})

	if err != nil {
		logger.Errorw("Failed to transfer player", "error", err)
		return newTransfer, err
	}

	return newTransfer, err
}

func GetTransferByPlayerID(playerID string) (model.Transfer, error) {
	var transfer model.Transfer

	if err := database.DB.First(&transfer, "player_id = ? AND transferred = ?", playerID, false).Error; err != nil {
		return transfer, err
	}

	return transfer, nil
}
