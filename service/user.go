package service

import (
	"errors"
	"strings"

	"github.com/nvzard/soccer-manager/database"
	"github.com/nvzard/soccer-manager/model"
	"github.com/nvzard/soccer-manager/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var logger *zap.SugaredLogger

func init() {
	logger = utils.GetLogger()
}

func CreateUser(user model.User) (model.User, error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		user.CreateTeam()
		result := tx.Create(&user)
		if err := result.Error; err != nil {
			if strings.Contains(err.Error(), model.UniqueConstraintEmail) {
				return errors.New("user with this email already exists")
			}

			logger.Errorw("Failed to create user", "error", err, "user", user)
			return err
		}

		// Update team_id for user
		user.TeamID = user.Team.ID
		result = tx.Save(&user)
		if err := result.Error; err != nil {
			logger.Errorw("Failed to update user's team", "error", err, "user", user)
			return err
		}
		return nil
	})

	if err != nil {
		logger.Errorw("Failed to create new user", "error", err)
		return user, err
	}

	return user, err
}

func GetUser(email string) (model.User, error) {
	var user model.User

	if err := database.DB.First(&user, "email = ?", email).Error; err != nil {
		return user, err
	}

	return user, nil
}
