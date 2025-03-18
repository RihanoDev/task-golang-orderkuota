package ports

import "orderkuota/domain"

type UserService interface {
	RegisterUser(userData *domain.User) (*domain.User, error)
	Login(email, password string) (*domain.User, string, error)
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id string) (*domain.User, error)
	UpdateUserByID(id string, userData *domain.User) (*domain.User, error)
	DeleteUserByID(id string) error
}
