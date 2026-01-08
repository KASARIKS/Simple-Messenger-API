package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kasariks/simple_messenger_api/config"
	"github.com/kasariks/simple_messenger_api/types"
	"github.com/kasariks/simple_messenger_api/utils"
)

const UserKey string = "userId"
const IatKey string = "iat"

func CreateJWT(secret []byte, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"iat":    time.Now().Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenString(r)
		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate the token: %v", err)
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userIdGotten := claims[UserKey].(string)
		iat := claims[IatKey]

		if time.Since(time.Unix(int64(iat.(float64)), 0)) > time.Duration(config.Envs.JWTExpirationInSeconds*int64(time.Second)) {
			log.Printf("the token is expired: %v", token)
			permissionDenied(w)
			return
		}

		u, err := store.GetUserById(userIdGotten)
		if err != nil {
			log.Printf("failed to get the user: %v", err)
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.Id)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func getTokenString(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	return tokenAuth
}

func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected string method: %v", t.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIdFromContext(ctx context.Context) string {
	userId, ok := ctx.Value(UserKey).(string)
	if !ok {
		return ""
	}

	return userId
}
