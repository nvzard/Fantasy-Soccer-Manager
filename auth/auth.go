package auth

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/nvzard/soccer-manager/model"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	TeamID uint   `json:"team_id"`
	jwt.StandardClaims
}

func GenerateJWT(user model.User) (tokenString string, err error) {
	// Token Expires after 1 hour
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		ID:     user.ID,
		Email:  user.Email,
		TeamID: user.TeamID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}
	return claims, nil
}

func GetUserAuth(context *gin.Context) (model.UserAuth, bool) {
	userAuth, exists := context.Get("user")
	return userAuth.(model.UserAuth), exists
}
