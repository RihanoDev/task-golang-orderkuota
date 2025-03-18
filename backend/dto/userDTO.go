package dto

import (
	"orderkuota/domain"
	"time"
)

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginResponse struct {
	User        *UserResponse `json:"user"`
	AccessToken string        `json:"access_token"`
}

func (u *UserRegister) ToDomain() *domain.User {
	return &domain.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u *UserLogin) ToDomain() *domain.User {
	return &domain.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u *UserUpdateRequest) ToDomain() *domain.User {
	user := &domain.User{}

	if u.Name != nil {
		user.Name = *u.Name
	}

	if u.Email != nil {
		user.Email = *u.Email
	}

	if u.Password != nil {
		user.Password = *u.Password
	}

	return user
}

func FromDomain(user *domain.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
