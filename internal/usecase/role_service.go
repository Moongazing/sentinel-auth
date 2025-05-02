package usecase

import (
	"SentinelAuth/internal/domain/role"
)

type RoleService struct {
	Repo role.Repository
}

func (s *RoleService) CreateRole(name string) (*role.Role, error) {
	r := &role.Role{Name: name}
	if err := s.Repo.Create(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (s *RoleService) GetAllRoles() ([]role.Role, error) {
	return s.Repo.GetAll()
}

func (s *RoleService) GetRoleByID(id uint) (*role.Role, error) {
	return s.Repo.GetByID(id)
}

func (s *RoleService) UpdateRole(id uint, name string) (*role.Role, error) {
	role, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	role.Name = name
	if err := s.Repo.Update(role); err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RoleService) DeleteRole(id uint) error {
	return s.Repo.Delete(id)
}
