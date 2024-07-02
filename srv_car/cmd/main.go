package main

import (
	"module/internal/config"
	"module/internal/database"
	"module/internal/server"

	log "github.com/sirupsen/logrus"
)

func main() {

	// установка логов. установка чтобы показывать логи debug уровня
	log.SetLevel(log.DebugLevel)
	log.Info("the server is starting")

	// получение конфигов
	cfg := config.ConfigMustLoad()

	// миграция и подключение к бд.
	database.OpenConnection(cfg.Connect)
	database.StartMigration()

	server.Consuming()

}
