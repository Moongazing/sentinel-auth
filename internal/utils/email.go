package utils

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func GenerateEmailToken() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func StoreEmailToken(rdb *redis.Client, token string, userID uint) error {
	key := fmt.Sprintf("emailverify:%s", token)
	return rdb.Set(context.Background(), key, userID, 24*time.Hour).Err()
}

func VerifyEmailToken(rdb *redis.Client, token string) (uint, error) {
	key := fmt.Sprintf("emailverify:%s", token)
	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}

	var userID uint
	_, _ = fmt.Sscanf(val, "%d", &userID)
	return userID, nil
}
