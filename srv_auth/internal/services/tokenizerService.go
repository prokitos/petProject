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

func TokenGetAccess(curUser models.Users) models.TokenResponser {
	var access string = createTokenAccess(curUser)

	responser := models.TokenResponser{
		AccessToken:  access,
		RefreshToken: "",
	}

	return responser
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
		log.Error("token dont signed")
		return ""
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
		log.Error("token dont signed")
		return ""
	}

	return tokenString
}

// метод для проведения проверки аксес токена
func TokenAccessValidate(bearer string) string {

	token, err := validateAccessToken(bearer)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "unathorized"
		}
		return "unathorized"
	}

	if !token.Valid {
		return "token expired"
	}

	// токен валиден. вернуть результат
	return "token useful"

}

// метод для проведения проверки рефреш токена
func TokenRefreshValidate(bearer string) string {

	token, err := validateAccessToken(bearer)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "unathorized"
		}
		return "unathorized"
	}

	if !token.Valid {
		return "token expired"
	}

	// токен валиден. вернуть результат
	return "token useful"

}

func GetAccessTokenFromRefresh(bearer string) string {

	token, err := validateRefreshToken(bearer)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "unathorized"
		}
		return "unathorized"
	}

	if !token.Valid {
		return "token expired"
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

	tokenString := strings.Split(bearerToken, " ")[1] // мы получаем токен в виде "bearer HG4HGK4FDRH45" и поэтому мы тут убираем слово bearer и пробел
	token, err := jwt.ParseWithClaims(tokenString, &models.RefreshToken{}, func(token *jwt.Token) (interface{}, error) {
		return RefreshKey, nil
	})

	return token, err
}
