package service

import (
	"strconv"

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

func TransferPlayer(player model.Player, transfer model.Transfer, buyingTeamID uint) error {
	sellingTeamID := player.TeamID
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// Reduce buyingTeam's AvailableCash
		buyingTeam, err := GetTeamByID(strconv.FormatUint(uint64(buyingTeamID), 10))
		if err != nil {
			logger.Errorw("Failed to fetch buyingTeam", "error", err, "buyingTeamID", buyingTeamID)
			return err
		}
		buyingTeam.AvailableCash -= transfer.AskedPrice
		result := tx.Save(&buyingTeam)
		if err := result.Error; err != nil {
			logger.Errorw("Failed to update team's available cash", "error", err, "buying_team", buyingTeam)
			return err
		}

		// Increase sellingTeam's AvailableCash
		sellingTeam, err := GetTeamByID(strconv.FormatUint(uint64(sellingTeamID), 10))
		if err != nil {
			logger.Errorw("Failed to fetch sellingTeam", "error", err, "sellingTeamID", sellingTeamID)
			return err
		}
		sellingTeam.AvailableCash += transfer.AskedPrice
		result = tx.Save(&sellingTeam)
		if err := result.Error; err != nil {
			logger.Errorw("Failed to update team's available cash", "error", err, "selling_team", sellingTeam)
			return err
		}

		// Remove from transfer list
		transfer.Transferred = true
		result = tx.Save(&transfer)
		if err := result.Error; err != nil {
			logger.Errorw("Failed to remove from transfer list", "error", err, "transfer", transfer)
			return err
		}

		// Update player's team and value
		player.Transfer(buyingTeamID)
		result = tx.Save(&player)
		if err := result.Error; err != nil {
			logger.Errorw("Failed to update team and player value", "error", err, "player", player)
			return err
		}

		return nil
	})

	if err != nil {
		logger.Errorw("Failed to transfer player", "error", err)
		return err
	}

	return err
}

func GetTransferByPlayerID(playerID string) (model.Transfer, error) {
	var transfer model.Transfer

	if err := database.DB.First(&transfer, "player_id = ? AND transferred = ?", playerID, false).Error; err != nil {
		return transfer, err
	}

	return transfer, nil
}

func GetAllPendingTransfers() ([]model.Transfer, error) {
	var transfers []model.Transfer

	if err := database.DB.Preload("Player").Find(&transfers, "transferred = ?", false).Error; err != nil {
		return transfers, err
	}

	return transfers, nil
}
