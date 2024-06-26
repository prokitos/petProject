package main

import (
	"module/internal/server"

	log "github.com/sirupsen/logrus"
)

func main() {

	// установка логов. установка чтобы показывать логи debug уровня
	log.SetLevel(log.DebugLevel)
	log.Info("the server is starting")

	server.Consuming()

}
