package services

import (
	"module/internal/models"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

var AccessKey []byte
var RefreshKey []byte

var AccessDuration time.Duration
var RefreshDuration time.Duration

// получение пары access и refresh token. передача refresh в базу данных
func TokenGetPair(curUser models.Users) models.TokenResponser {

	var access string = createTokenAccess(curUser)
	var refresh string = createTokenRefresh(curUser, access)

	responser := models.TokenResponser{
		AccessToken:  access,
		RefreshToken: refresh,
	}

	return responser
}

func GetUserFromRefresh(refresh string, access string) (models.Users, error) {

	// проверка рефреш токена, валидность, попытка расшифровать его
	token, err := validateRefreshToken(refresh)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Debug("token dont read")
			return models.Users{}, models.ResponseTokenDontRead()
		}
		log.Debug("token is expired")
		return models.Users{}, models.ResponseTokenExpired()
	}

	// рефреш токен не валиден, неизвестная ошибка!!
	if !token.Valid {
		log.Debug("invalid token")
		return models.Users{}, models.ResponseTokenError()
	}

	// токен валиден. вернуть новый аксес токен
	user := token.Claims.(*models.RefreshToken)

	// аксес токен не подходит к рефреш токену
	accessed := strings.Split(access, " ")[1]
	if user.AcceessToken != accessed {
		log.Debug("access and refresh token dont same")
		return models.Users{}, models.ResponseTokenDontSame()
	}

	var temp models.Users
	temp.Login = user.Login
	temp.AccessLevel = user.AccessLevel

	return temp, nil
}

// создание аксес токена.
func createTokenAccess(curUser models.Users) string {

	// создаем токен
	var tokenObj = models.AccessToken{
		Login:       curUser.Login,
		AccessLevel: curUser.AccessLevel,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessDuration).Unix(),
		},
	}

	// шифруем с помощью accessKey. метод HS = HMAC + SHA 512
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenObj)
	tokenString, err := token.SignedString(AccessKey)
	if err != nil {
		log.Debug("token dont sighed")
		return models.ResponseTokenError().Error()
	}

	// возвращаем
	return tokenString
}

// создание рефреш токена. срок жини 5 минут для теста
func createTokenRefresh(curUser models.Users, accessToken string) string {

	// создаем токен
	var tokenObj = models.RefreshToken{
		Login:        curUser.Login,
		AccessLevel:  curUser.AccessLevel,
		AcceessToken: accessToken,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(RefreshDuration).Unix(),
		},
	}

	// шифруем с помощью refreshKey. метод HS = HMAC + SHA 512
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenObj)
	tokenString, err := token.SignedString(RefreshKey)
	if err != nil {
		log.Debug("token dont signed")
		return models.ResponseTokenError().Error()
	}

	return tokenString
}

// метод для проведения проверки аксес токена
func TokenAccessValidate(bearer string) (string, int) {

	token, err := validateAccessToken(bearer)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Debug("token dont read")
			return models.ResponseTokenUnauthorized().Error(), 0
		}
		log.Debug("token expired")
		return models.ResponseTokenUnauthorized().Error(), 0
	}

	if !token.Valid {
		log.Debug("invalid token")
		return models.ResponseTokenExpired().Error(), 0
	}

	user := token.Claims.(*models.AccessToken)

	// токен валиден. вернуть результат
	return models.ResponseTokenGood().Error(), user.AccessLevel

}

// метод для проведения проверки рефреш токена
func TokenRefreshValidate(bearer string) string {

	token, err := validateRefreshToken(bearer)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Debug("token dont read")
			return models.ResponseTokenUnauthorized().Error()
		}
		log.Debug("token expired")
		return models.ResponseTokenUnauthorized().Error()
	}

	if !token.Valid {
		log.Debug("invalid token")
		return models.ResponseTokenExpired().Error()
	}

	// токен валиден. вернуть результат
	return models.ResponseTokenGood().Error()

}

func GetAccessTokenFromRefresh(bearer string) string {

	token, err := validateRefreshToken(bearer)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Debug("token dont read")
			return models.ResponseTokenUnauthorized().Error()
		}
		log.Debug("token expired")
		return models.ResponseTokenUnauthorized().Error()
	}

	if !token.Valid {
		log.Debug("invalid token")
		return models.ResponseTokenExpired().Error()
	}

	// токен валиден. вернуть новый аксес токен
	user := token.Claims.(*models.RefreshToken)

	var temp models.Users
	temp.Login = user.Login
	temp.AccessLevel = user.AccessLevel

	result := createTokenAccess(temp)

	return result
}

// проверка валидности access токена
func validateAccessToken(bearerToken string) (*jwt.Token, error) {

	tokenString := strings.Split(bearerToken, " ")[1] // мы получаем токен в виде "bearer HG4HGK4FDRH45" и поэтому мы тут убираем слово bearer и пробел
	token, err := jwt.ParseWithClaims(tokenString, &models.AccessToken{}, func(token *jwt.Token) (interface{}, error) {
		return AccessKey, nil
	})
	return token, err
}

// проверка валидности refresh токена
func validateRefreshToken(bearerToken string) (*jwt.Token, error) {

	tokenString := strings.Split(bearerToken, " ")[0]
	token, err := jwt.ParseWithClaims(tokenString, &models.RefreshToken{}, func(token *jwt.Token) (interface{}, error) {
		return RefreshKey, nil
	})

	return token, err
}
