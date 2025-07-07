package main

import (
	"go-29/cmd"
	"go-29/internal/data/repository"
	"go-29/internal/wire"
	"go-29/pkg/database"
	"go-29/pkg/middleware"
	"go-29/pkg/utils"
	"log"

	"go.uber.org/zap"
)

func main() {
	// read config
	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	// init logger
	logger, err := utils.InitLogger(config.PathLogger, config)
	if err != nil {
		log.Fatal("can't init logger %w", zap.Error(err))
	}

	//Init db
	db, err := database.InitDB(config)
	if err != nil {
		logger.Fatal("can't connect to database ", zap.Error(err))
	}

	repo := repository.NewRepository(db, logger)
	mLogger := middleware.NewLoggerMiddleware(logger)
	router := wire.Wiring(repo, mLogger, logger, config)

	cmd.ApiServer(config, logger, router)
}
