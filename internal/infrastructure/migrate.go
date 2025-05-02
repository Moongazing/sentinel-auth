package infrastructure

import (
	"SentinelAuth/internal/domain/role"
	"SentinelAuth/internal/domain/user"
	"log"
)

func MigrateDB() {
	err := DB.AutoMigrate(&role.Role{}, &user.User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
