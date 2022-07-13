package service

// func CreateTransfer(newTransfer model.Transfer) (model.Transfer, error) {
// 	err := database.DB.Transaction(func(tx *gorm.DB) error {
// 		result := tx.Create(&newTransfer)
// 		if err := result.Error; err != nil {
// 			if strings.Contains(err.Error(), model.UniqueConstraintEmail) {
// 				return errors.New("user with this email already exists")
// 			}

// 			logger.Errorw("Failed to create user", "error", err, "user", user)
// 			return err
// 		}

// 		// Update team_id for user
// 		user.TeamID = user.Team.ID
// 		result = tx.Save(&user)
// 		if err := result.Error; err != nil {
// 			logger.Errorw("Failed to update user's team", "error", err, "user", user)
// 			return err
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		logger.Errorw("Failed to create new user", "error", err)
// 		return user, err
// 	}

// 	return user, err
// }
