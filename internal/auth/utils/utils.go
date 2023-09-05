package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type TokenType int32

const (
	TokenTypeAccess TokenType = iota
	TokenTypeRefresh
)

// CreatJwt 生成一个JwtToken，包含uid
func CreatJwt(uid string, tp TokenType, key string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  uid,
		"type": tp,
		"exp":  time.Now().Add(6 * time.Hour).Unix(),
	})
	token, err := at.SignedString([]byte(key))
	if err != nil {
		return "", ErrSignedString
	}
	return token, nil
}

// ParseToken 从Jwt中解析Token
func ParseToken(token string, tokenType TokenType, key string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenSigningMethod
		}
		return []byte(key), nil
	})
	var ve *jwt.ValidationError
	if errors.As(err, &ve) {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", ErrTokenMalformed
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return "", ErrTokenExpired
		} else {
			return "", ErrTokenHandle
		}
	}
	if claims, ok := claim.Claims.(jwt.MapClaims); ok && claim.Valid {
		if tp, ok := claims["type"]; ok && tp.(float64) == float64(tokenType) {
			return claims["uid"].(string), nil
		}
	}
	return "", ErrTokenHandle
}

func GenerateUUID() string {
	return uuid.New().String()
}
