package utils

import (
	"errors"
	"time"

	"server_zzq/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT 声明
type Claims struct {
	ShopID uint `json:"shop_id"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT token
func GenerateToken(shopID uint) (string, error) {
	cfg := config.GetConfig()
	if cfg == nil {
		return "", errors.New("config not initialized")
	}

	claims := Claims{
		ShopID: shopID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.JWT.Expire) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "server_zzq",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}

// ParseToken 解析 JWT token
func ParseToken(tokenString string) (*Claims, error) {
	cfg := config.GetConfig()
	if cfg == nil {
		return nil, errors.New("config not initialized")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
