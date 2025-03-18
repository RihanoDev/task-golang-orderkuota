package ports

import "orderkuota/domain"

type UserRepository interface {
	CreateUser(userData *domain.User) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	UpdateUserByID(id string, userData *domain.User) (*domain.User, error)
	DeleteUserByID(id string) error
}
