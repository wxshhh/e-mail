package utils

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

var jwtSecret = []byte("wxshhh")

type Claims struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Authority string `json:"authority"`
	jwt.StandardClaims
}

type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

// GenerateToken 签发token
func GenerateToken(id uint, username string, authority string) (string, error) {
	now := time.Now()
	expireTime := now.Add(24 * time.Hour)
	claims := Claims{
		ID:        id,
		Username:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "wxshhh",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	// token: Bearer xxxx.xxxxxx.xxxx
	tokenStr := strings.Split(token, " ")[1]
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(tokenStr *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// GenerateEmailToken 生成邮箱token
func GenerateEmailToken(userID uint, email string, password string, operationType uint) (string, error) {
	now := time.Now()
	expireTime := now.Add(24 * time.Hour)
	claims := EmailClaims{
		UserID:        userID,
		Email:         email,
		Password:      password,
		OperationType: operationType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "wxshhh",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseEmailToken 验证邮箱token
func ParseEmailToken(token string) (*EmailClaims, error) {
	// token: Bearer xxxx.xxxxxx.xxxx
	tokenStr := strings.Split(token, " ")[1]
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &EmailClaims{}, func(tokenStr *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
