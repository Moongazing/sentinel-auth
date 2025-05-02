package utils

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/redis/go-redis/v9"
)

func GenerateOTP() (string, error) {
	max := big.NewInt(1000000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", n.Int64()), nil
}

func StoreOTP(rdb *redis.Client, userID uint, otp string) error {
	key := fmt.Sprintf("otp:%d", userID)
	return rdb.Set(context.Background(), key, otp, 5*time.Minute).Err()
}

func VerifyOTP(rdb *redis.Client, userID uint, otp string) bool {
	key := fmt.Sprintf("otp:%d", userID)
	val, err := rdb.Get(context.Background(), key).Result()
	return err == nil && val == otp
}
