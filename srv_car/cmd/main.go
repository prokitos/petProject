package main

import (
	"module/internal/config"
	"module/internal/database"
	"module/internal/server"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {

	// установка логов. установка чтобы показывать логи debug уровня
	log.SetLevel(log.DebugLevel)
	log.Info("the server is starting")

	// получение конфигов
	cfg := config.ConfigMustLoad("docker")
	server.RMQaddress = cfg.External.RabbitMqServer

	// проверка что есть бд не делаю тут, так как есть в сервисе авторизации, а без авторизации этот сервис не работает

	// миграция и подключение к бд.
	database.OpenConnection(cfg.Connect)
	database.StartMigration()

	// запуск rabbitMQ в потоках
	go server.CarSellConsuming()
	go server.CarConsuming()

	// завершение при нажатии кнопок
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

}
