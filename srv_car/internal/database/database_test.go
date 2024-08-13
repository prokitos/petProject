package database

import (
	"module/internal/models"
	"testing"
)

// конфиги для доступа к тестовой бд
func getTestConfig() models.ConnectConfig {

	var cfg models.ConnectConfig
	cfg.Host = "localhost"
	cfg.Name = "pets_test"
	cfg.Pass = "root"
	cfg.Port = "8092"
	cfg.User = "postgres"
	cfg.Reload = false

	return cfg
}

// проверка подключения к базе
func TestConnetion(t *testing.T) {

	cfg := getTestConfig()
	CheckDatabaseCreated(cfg)
	OpenConnection(cfg)

	StartMigration()

	sqlDB, err := GlobalHandler.DB()
	if err != nil {
		t.Errorf(err.Error())
	}

	qq := sqlDB.Stats()
	if qq.Idle == 0 && qq.InUse == 0 {
		t.Errorf("result wrong at test, does not connect to server")
	}

	sqlDB.Close()

}

// проверка что машина создается
func TestCreateCar(t *testing.T) {

	cfg := getTestConfig()
	OpenConnection(cfg)

	res := CreateNewCar(models.Car{Mark: "Toyota", Year: "1999", Price: 1000000})

	if res.Code != models.ResponseCarGoodCreate().Code {
		t.Errorf("car don't create")
	}

}

// попытка изменить марку
func TestUpdateCar(t *testing.T) {

	cfg := getTestConfig()
	OpenConnection(cfg)

	res := UpdateCar(models.Car{Id: 1, Mark: "Lada"})

	if res.Code != models.ResponseCarGoodUpdate().Code {
		t.Errorf("database update car error")
	}

}

// проверка что машина есть в бд
func TestCheckCar(t *testing.T) {

	cfg := getTestConfig()
	OpenConnection(cfg)

	res := ShowCar(models.Car{Mark: "Lada"})

	if res.Code == models.ResponseCarBadShow().Code {
		t.Errorf("database don't has a car")
	}

}

// попытка удалить машину
func TestCarDelete(t *testing.T) {

	cfg := getTestConfig()
	OpenConnection(cfg)

	res := DeleteCar(models.Car{Id: 1})

	if res.Code != models.ResponseCarGoodDelete().Code {
		t.Errorf("database delete car error")
	}

}
