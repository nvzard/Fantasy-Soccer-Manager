package main

import (
	"github.com/nvzard/soccer-manager/database"
	"github.com/nvzard/soccer-manager/server"
	"github.com/nvzard/soccer-manager/utils"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	logger = utils.GetLogger()
}

func main() {
	// // Connect to Database
	err := database.Connect()
	if err != nil {
		logger.Fatalw("Failed to connect to database", "error", err.Error())
	}

	logger.Info("Soccer Manager Service started")

	// Prepare Database
	err = database.Prepare()
	if err != nil {
		logger.Fatalw("Failed to prepare database", "error", err.Error())
	}

	// Start Webserver
	r := server.SetupApiServer()
	err = r.Run()
	if err != nil {
		logger.Fatalw("Failed to start server", "error", err.Error())
	}
}
