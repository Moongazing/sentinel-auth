package repository

import (
	"SentinelAuth/internal/domain/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *user.User) error {
	return r.DB.Create(user).Error
}
func (r *UserRepository) GetByEmail(email string) (*user.User, error) {
	var u user.User
	if err := r.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetByID(id uint) (*user.User, error) {
	var u user.User
	if err := r.DB.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetPaginated(offset, limit int) ([]user.User, int64, error) {
	var users []user.User
	var total int64

	if err := r.DB.Model(&user.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Offset(offset).Limit(limit).Order("created_at desc").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
func (r *UserRepository) UpdateRole(userID uint, roleID uint) error {
	return r.DB.Model(&user.User{}).
		Where("id = ?", userID).
		Update("role_id", roleID).Error
}
func (r *UserRepository) MarkEmailVerified(userID uint) error {
	return r.DB.Model(&user.User{}).
		Where("id = ?", userID).
		Update("is_email_verified", true).Error
}
