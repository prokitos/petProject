package database

import (
	"fmt"
	"module/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

var GlobalHandler *gorm.DB

// открыть соединение, и поместить его в глобальную переменну.
func OpenConnection(config models.ConnectConfig) {
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Pass, config.Host, config.Port, config.Name)

	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	GlobalHandler = db
}

// миграция
func StartMigration() {

	GlobalHandler.AutoMigrate(models.CarEngine{})
	GlobalHandler.AutoMigrate(models.People{})
	GlobalHandler.AutoMigrate(models.AdditionalDevices{})
	GlobalHandler.AutoMigrate(models.Car{})
	GlobalHandler.AutoMigrate(models.Selling{})

	log.Info("migration complete")

}
