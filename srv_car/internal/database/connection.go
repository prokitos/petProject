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
		panic("connection with postgres error")
		return
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

func CheckDatabaseCreated(config models.ConnectConfig) error {

	// открытие соеднение с базой по стандарту
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Pass, config.Host, config.Port, "postgres")
	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		log.Error("database don't open")
		return models.ResponseBadRequest()
	}

	// закрытие бд
	sql, _ := db.DB()
	defer func() {
		_ = sql.Close()
	}()

	// проверка если есть нужная нам база данных
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", config.Name)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		log.Error("error, dont read bd")
		return models.ResponseBadRequest()
	}

	// если нет, то создать
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", config.Name)
		if rs := db.Exec(stmt); rs.Error != nil {
			log.Error("error, dont create a database")
			return models.ResponseBadRequest()
		}
	}

	return nil
}
