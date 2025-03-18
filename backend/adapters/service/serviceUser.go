package service

import (
	"fmt"
	"orderkuota/domain"
	ports "orderkuota/ports/repository"
	"orderkuota/utils"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceAdapter struct {
	repository ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserServiceAdapter {
	return &UserServiceAdapter{repository: repo}
}

func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func (us *UserServiceAdapter) RegisterUser(userData *domain.User) (*domain.User, error) {
	if userData.Email == "" || userData.Password == "" {
		return nil, fmt.Errorf("email and password can't be empty")
	}

	if !isValidEmail(userData.Email) {
		return nil, fmt.Errorf("invalid email format")
	}

	existingUserEmail, _ := us.repository.GetUserByEmail(userData.Email)
	if existingUserEmail != nil {
		return nil, fmt.Errorf("email is already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	userData.ID = uuid.New().String()
	userData.Password = string(hashedPassword)
	userData.CreatedAt = time.Now()
	userData.UpdatedAt = time.Now()

	user, err := us.repository.CreateUser(userData)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

func (us *UserServiceAdapter) Login(email, password string) (*domain.User, string, error) {
	if email == "" || password == "" {
		return nil, "", fmt.Errorf("email and password cannot be empty")
	}

	user, err := us.repository.GetUserByEmail(email)
	if err != nil {
		return nil, "", fmt.Errorf("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", fmt.Errorf("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, "", fmt.Errorf("Invalid credentials")
	}

	user.Password = ""
	return user, token, nil
}

func (us *UserServiceAdapter) GetAllUsers() ([]domain.User, error) {
	users, err := us.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// for i := range users {
	// 	users[i].Password = ""
	// }

	return users, nil
}

func (us *UserServiceAdapter) GetUserByID(id string) (*domain.User, error) {
	user, err := us.repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// user.Password = ""
	return user, nil
}

func (us *UserServiceAdapter) UpdateUserByID(id string, userData *domain.User) (*domain.User, error) {
	existingUser, err := us.repository.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	if userData.Name != "" {
		existingUser.Name = userData.Name
	}
	if userData.Email != "" {
		if !isValidEmail(userData.Email) {
			return nil, fmt.Errorf("invalid email format")
		}
		existingUser.Email = userData.Email
	}
	if userData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("failed to hash password: %v", err)
		}
		userData.Password = string(hashedPassword)
	} else {
		userData.Password = existingUser.Password
	}

	userData.UpdatedAt = time.Now()
	updatedUser, err := us.repository.UpdateUserByID(id, userData)
	if err != nil {
		return nil, err
	}

	// updatedUser.Password = ""
	return updatedUser, nil
}

func (us *UserServiceAdapter) DeleteUserByID(id string) error {
	_, err := us.repository.GetUserByID(id)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	return us.repository.DeleteUserByID(id)
}
