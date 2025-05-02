package repository

import (
	"SentinelAuth/internal/domain/role"
	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (r *RoleRepository) Create(role *role.Role) error {
	return r.DB.Create(role).Error
}

func (r *RoleRepository) GetAll() ([]role.Role, error) {
	var roles []role.Role
	err := r.DB.Find(&roles).Error
	return roles, err
}

func (r *RoleRepository) GetByID(id uint) (*role.Role, error) {
	var role role.Role
	err := r.DB.First(&role, id).Error
	return &role, err
}

func (r *RoleRepository) Update(role *role.Role) error {
	return r.DB.Save(role).Error
}

func (r *RoleRepository) Delete(id uint) error {
	return r.DB.Delete(&role.Role{}, id).Error
}
