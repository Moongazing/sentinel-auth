package user

type Repository interface {
	Create(user *User) error

	GetByEmail(email string) (*User, error)

	GetByID(id uint) (*User, error)

	GetPaginated(offset, limit int) ([]User, int64, error)
	UpdateRole(userID uint, roleID uint) error
	MarkEmailVerified(userID uint) error
}
