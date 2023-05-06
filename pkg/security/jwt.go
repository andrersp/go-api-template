package secutiry

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/andrersp/go-api-template/internal/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func CreateToken(ID uuid.UUID) (accessToken string, err error) {

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

	accessToken, err = claims.SignedString([]byte(config.SECRET_TOKEN))
	return
}

func GetTokenData(ctx context.Context) (TokenData, error) {

	contextValue := ctx.Value(TokenData{})

	if contextValue == nil {
		err := errors.New("context not found")

		return TokenData{}, err
	}

	tokenData := contextValue.(TokenData)
	return tokenData, nil

}

func ExtractTokenData(r *http.Request) (tokenData TokenData, err error) {

	tokenString, err := extractTokenString(r)

	if err != nil {
		return
	}

	token, err := verifyToken(tokenString)
	if err != nil {
		return
	}

	if err = tokenValid(token); err != nil {
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

func extractTokenString(r *http.Request) (tokenString string, err error) {

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

func verifyToken(tokenString string) (*jwt.Token, error) {

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
	fmt.Println(config.SECRET_TOKEN)

	return []byte(config.SECRET_TOKEN), nil

}

func tokenValid(token *jwt.Token) (err error) {

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return ErrUnauthorized
}
