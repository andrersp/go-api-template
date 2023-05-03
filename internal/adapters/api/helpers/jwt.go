package helpers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/andrersp/go-api-template/internal/config"
	"github.com/andrersp/go-api-template/internal/core/dto"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var (
	ErrUnauthorized = errors.New("UNAUTHORIZED")
	ErrForbiden     = errors.New("FORBIDEN")
)

type jwtLoginClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func CreateToken(ID uuid.UUID) (loginResponse dto.LoginResponse, err error) {

	now := time.Now()
	atClaims := jwtLoginClaims{
		ID: ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
			Subject:   "Sub",
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := claims.SignedString([]byte("4a031514c89b259aaad22c00bbd0be629412ea8d3926f39da330e89933065214"))
	if err != nil {
		return
	}
	loginResponse.AccessToken = accessToken
	loginResponse.TokenType = "Bearer"
	return

}

func ExtractTokenData(r *http.Request) (tokenData dto.TokenData, err error) {

	tokenString, err := ExtractTokenString(r)

	if err != nil {
		return
	}

	token, err := VerifyToken(tokenString)
	if err != nil {
		return
	}

	if err = TokenValid(token); err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimsID := fmt.Sprintf("%s", claims["id"])

		tokenData.UserID = claimsID

	} else {
		err = ErrForbiden
	}

	return

}

func ExtractTokenString(r *http.Request) (tokenString string, err error) {

	if _, ok := r.Header["Authorization"]; !ok {
		err = ErrUnauthorized
		return
	}
	bearerToken := r.Header.Get("Authorization")

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		err = ErrUnauthorized
		return
	}
	tokenString = splitToken[1]
	return
}

func VerifyToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, verifyTokenKey)
	if err != nil {
		return nil, ErrUnauthorized
	}

	return token, nil
}

func verifyTokenKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, ErrUnauthorized
	}

	return []byte(config.SECRET_TOKEN), nil

}

func TokenValid(token *jwt.Token) (err error) {

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return ErrUnauthorized
}
