package main

import (
	"module/internal/app"

	log "github.com/sirupsen/logrus"
)

func main() {

	// установка логов. установка чтобы показывать логи debug уровня
	log.SetLevel(log.DebugLevel)
	log.Info("the server is starting")

	var application app.App
	application.NewServer(":8003")

}
