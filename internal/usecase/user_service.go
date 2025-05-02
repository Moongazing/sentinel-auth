package usecase

import (
	"SentinelAuth/internal/domain/user"
	"SentinelAuth/internal/utils"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type UserService struct {
	Repo user.Repository
	Rdb  *redis.Client
}

func (s *UserService) Register(email, password, fullName string) (*user.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser := &user.User{
		Email:        email,
		PasswordHash: hashedPassword,
		FullName:     fullName,
	}

	if err := s.Repo.Create(newUser); err != nil {
		return nil, err
	}
	token, err := utils.GenerateEmailToken()
	if err == nil {
		_ = utils.StoreEmailToken(s.Rdb, token, newUser.ID)
		fmt.Printf("Email verification token for user %s: http://localhost:8080/api/verify-email?token=%s\n", newUser.Email, token)
	}

	return newUser, nil
}

func (s *UserService) Login(email, password string) (string, string, error) {
	u, err := s.Repo.GetByEmail(email)
	if err != nil {
		return "", "", err
	}

	if !utils.CheckPasswordHash(password, u.PasswordHash) {
		return "", "", errors.New("invalid credentials")
	}

	access, err := utils.GenerateAccessToken(u.ID, u.Email)
	if err != nil {
		return "", "", err
	}

	refresh, err := utils.GenerateRefreshToken(u.ID)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (s *UserService) RefreshToken(refreshToken string) (string, error) {
	isBlacklisted, _ := utils.IsRefreshTokenBlacklisted(s.Rdb, refreshToken)
	if isBlacklisted {
		return "", errors.New("refresh token is revoked")
	}
	userID, err := utils.ParseRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	user, err := s.Repo.GetByID(userID)
	if err != nil {
		return "", err
	}

	access, err := utils.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return access, nil
}

func (s *UserService) GetUsers(page, pageSize int) ([]user.User, int64, error) {
	offset := (page - 1) * pageSize
	return s.Repo.GetPaginated(offset, pageSize)
}
func (s *UserService) AssignRoleToUser(userID, roleID uint) error {
	return s.Repo.UpdateRole(userID, roleID)
}

func (s *UserService) MarkEmailVerified(userID uint) error {
	return s.Repo.MarkEmailVerified(userID)
}
