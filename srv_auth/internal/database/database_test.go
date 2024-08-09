package database

import (
	"fmt"
	"module/internal/models"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

// проверка что пользователь создаётся
func TestCreateUser(t *testing.T) {

	cfg := getTestConfig()
	OpenConnection(cfg)

	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	_, err := CreateNewUser(c, models.Users{Login: "Test", Password: "Test", AccessLevel: 1, RefreshToken: "test"})

	if err != nil {
		t.Errorf("user don't create")
	}

}

// проверка что пользователь есть в бд
func TestCheckUser(t *testing.T) {

	cfg := getTestConfig()
	OpenConnection(cfg)

	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	err := CheckUserName(c, models.Users{Login: "Test"})

	if err == nil || err.Error() != models.ResponseBadRequest().Error() {
		t.Errorf("database don't has a user")
	}

}

// попытка увеличить уровень доступа
func TestUpdateLevel(t *testing.T) {

	cfg := getTestConfig()
	OpenConnection(cfg)

	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	err := UpdateUserLevel(c, models.Users{Login: "Test"})

	if err != nil {
		t.Errorf("database update level error")
	}

}

// попытка изменить токен пользователя
func TestUpdateToken(t *testing.T) {

	cfg := getTestConfig()
	OpenConnection(cfg)

	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	err := UpdateRefreshToken(c, models.Users{Login: "Test", RefreshToken: "newToken"})

	if err != nil {
		t.Errorf("database update token error")
	}

}

// проверка что уровень и пароль изменились
func TestGetUser(t *testing.T) {

	cfg := getTestConfig()
	OpenConnection(cfg)

	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	user, err := GetExistingUser(c, models.Users{Login: "Test"})
	resultUser := models.Users{Login: "Test", Password: "Test", AccessLevel: 5, RefreshToken: "newToken"}

	if user.AccessLevel != resultUser.AccessLevel || user.Login != resultUser.Login || user.RefreshToken != resultUser.RefreshToken || err != nil {
		t.Errorf("database get user error")
	}

}

// удаляем тестовую базу данных после тестов.
func TestEnd(t *testing.T) {

	cfg := getTestConfig()

	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Pass, cfg.Host, cfg.Port, "postgres")
	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		t.Errorf("database dont open")
	}

	// закрытие бд
	sql, _ := db.DB()
	defer func() {
		_ = sql.Close()
	}()

	stmt := fmt.Sprintf("DROP DATABASE %s;", cfg.Name)
	if rs := db.Exec(stmt); rs.Error != nil {
		t.Errorf("database dont delete")
	}

}
