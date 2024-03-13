package main

import (
	"log"
	"os"

	"test_service_filmoteka/config"
	"test_service_filmoteka/internal/server"
	"test_service_filmoteka/pkg/db/postgres"
	"test_service_filmoteka/pkg/logger"
	"test_service_filmoteka/pkg/utils"
)

// @title Go app
// @version 1.0
// @description Golang app
// @contact.name Alimadad Ismoilov
// @contact.url https://github.com/AliIsmoilov
// @contact.email alimadadismoilov@gmail.com
// @BasePath /api/v1
func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	}
	// defer psqlDB.Close()

	s := server.NewServer(cfg, psqlDB, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
