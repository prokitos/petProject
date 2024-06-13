package main

import (
	"module/internal/app"
	"module/internal/config"
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
	cfg := config.ConfigMustLoad()
	var application app.App

	// запуск сервера в горутине, чтобы потом нормально звершать приложение
	go application.NewServer(cfg.Server)

	// в итоге мы обрабатываем завершение приложения, и если мы закрыаем его как либо, то оно выполняет действие из метода Stop
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Stop() // безопасное выключение сервера

}
