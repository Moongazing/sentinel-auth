package role

type Repository interface {
	Create(role *Role) error
	GetAll() ([]Role, error)
	GetByID(id uint) (*Role, error)
	Update(role *Role) error
	Delete(id uint) error
}
