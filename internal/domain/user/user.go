package user

import "time"

type User struct {
	ID              uint   `gorm:"primaryKey"`
	Email           string `gorm:"uniqueIndex;not null"`
	PasswordHash    string `gorm:"not null"`
	FullName        string
	IsEmailVerified bool   `gorm:"default:false"`
	IsOTPEnabled    bool   `gorm:"default:false"`
	RoleID          *uint  `gorm:"index"`
	RoleName        string `gorm:"-"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
