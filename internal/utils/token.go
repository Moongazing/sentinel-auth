package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

func GenerateAccessToken(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseRefreshToken(tokenStr string) (uint, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["sub"] == nil {
		return 0, errors.New("invalid claims")
	}

	userID := uint(claims["sub"].(float64)) // JWT numeric types are float64
	return userID, nil
}

func BlacklistRefreshToken(rdb *redis.Client, token string, exp time.Duration) error {
	key := fmt.Sprintf("blacklist:%s", token)
	return rdb.Set(context.Background(), key, true, exp).Err()
}

func IsRefreshTokenBlacklisted(rdb *redis.Client, token string) (bool, error) {
	key := fmt.Sprintf("blacklist:%s", token)
	return rdb.Exists(context.Background(), key).Val() == 1, nil
}
