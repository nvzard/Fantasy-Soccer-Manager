package controller

// func CreateTransfer(context *gin.Context) {
// 	var newTransfer model.Transfer
// 	if err := context.ShouldBindJSON(&newTransfer); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		context.Abort()
// 		return
// 	}

// 	player, err := service.GetPlayerByID(string(newTransfer.PlayerID))
// 	if err != nil {
// 		context.JSON(http.StatusNotFound, gin.H{"error": "player does not exist"})
// 		context.Abort()
// 		return
// 	}

// 	transfer, err := service.CreateTransfer(newTransfer, player)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		context.Abort()
// 		return
// 	}
// 	context.JSON(http.StatusCreated, gin.H{"transfer_id": transfer.ID})
// }

// func GetUser(context *gin.Context) {
// 	userEmail := context.Params.ByName("email")

// 	user, err := service.GetUser(userEmail)

// 	if err != nil {
// 		context.JSON(http.StatusNotFound, gin.H{"error": "user does not exist"})
// 		context.Abort()
// 		return
// 	}

// 	context.JSON(http.StatusCreated, gin.H{
// 		"id":         user.ID,
// 		"email":      user.Email,
// 		"first_name": user.FirstName,
// 		"last_name":  user.LastName,
// 		"team_id":    user.TeamID,
// 	})
// }
