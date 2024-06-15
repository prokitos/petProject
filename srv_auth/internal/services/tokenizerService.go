package services

import (
	"module/internal/models"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

var accessKey = []byte("basic_key")
var refreshKey = []byte("super_mega_key")

var accessDuration time.Duration = time.Minute
var refreshDuration time.Duration = time.Minute * 5

// получение пары access и refresh token. передача refresh в базу данных
func TokenGetPair(curUser models.Users) models.TokenResponser {

	var access string = createTokenAccess(curUser)
	var refresh string = createTokenRefresh(curUser, access)

	responser := models.TokenResponser{
		AccessToken:  access,
		RefreshToken: refresh,
	}

	// запись рефреш токена в базу данных, значит потом его проверять при новом рефреш токене

	return responser
}

// создание аксес токена.
func createTokenAccess(curUser models.Users) string {

	// создаем токен
	var tokenObj = models.AccessToken{
		Login:       curUser.Login,
		AccessLevel: curUser.AccessLevel,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessDuration).Unix(),
		},
	}

	// шифруем с помощью accessKey. метод HS = HMAC + SHA 512
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenObj)
	tokenString, err := token.SignedString(accessKey)
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
			ExpiresAt: time.Now().Add(refreshDuration).Unix(),
		},
	}

	// шифруем с помощью refreshKey. метод HS = HMAC + SHA 512
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenObj)
	tokenString, err := token.SignedString(refreshKey)
	if err != nil {
		log.Error("token dont signed")
		return ""
	}

	return tokenString
}

// проверка валидности access токена
func validateAccessToken(bearerToken string) (*jwt.Token, error) {

	tokenString := strings.Split(bearerToken, " ")[1]
	token, err := jwt.ParseWithClaims(tokenString, &models.AccessToken{}, func(token *jwt.Token) (interface{}, error) {
		return accessKey, nil
	})
	return token, err
}

// проверка валидности refresh токена
func validateRefreshToken(bearerToken string) (*jwt.Token, error) {

	tokenString := bearerToken
	token, err := jwt.ParseWithClaims(tokenString, &models.RefreshToken{}, func(token *jwt.Token) (interface{}, error) {
		return refreshKey, nil
	})

	return token, err
}

// // получить новый рефреш токен
// func RenewToken(refreshToken string, accessToken string) models.TokenResponser {

// 	var result models.TokenResponser

// 	// проверка рефреш токена
// 	token, err := validateRefreshToken(refreshToken)
// 	if err != nil {
// 		result.RefreshToken = "refresh token unauthorized"
// 		if err == jwt.ErrSignatureInvalid {
// 			result.RefreshToken = "refresh token sign unknown"
// 			return result
// 		}
// 		return result
// 	}

// 	if !token.Valid {
// 		result.RefreshToken = "refresh token expired"
// 		return result
// 	}

// 	refToken := token.Claims.(*TokenRefreshData)

// 	if refToken.AcceessToken != accessToken {
// 		result.RefreshToken = "access token missmatch"
// 		return result
// 	}

// 	// токен валиден. удаляем рефреш токен из базы, получаем новый аксес токен
// 	var DBmodel models.TokenDB
// 	DBmodel.GUID = refToken.GUID
// 	DBmodel.RefreshToken = refreshToken

// 	// проверка что в базе есть такой токен, и что он принадлежит этому пользователю
// 	retModel := database.SearchToken(DBmodel)
// 	if retModel.RefreshToken != DBmodel.RefreshToken {
// 		result.RefreshToken = "wrong user in token"
// 		return result
// 	}

// 	// удаление всех токенов (если как-то получилось много) у данного GUID из базы (возможно не надо)
// 	// !!!!!!!!!!!!!!!!!!!!!!
// 	var DBmodelDel models.TokenDB
// 	DBmodelDel.GUID = refToken.GUID
// 	database.DeleteToken(DBmodelDel)
// 	// !!!!!!!!!!!!!!!!!!!!!!

// 	// создание нового рефреш и аксес токена
// 	result = TokenGetPair(refToken.GUID)

// 	return result
// }
