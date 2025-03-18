package repository

import (
	"database/sql"
	"fmt"
	"orderkuota/domain"
	ports "orderkuota/ports/repository"
)

type UserRepositoryAdapter struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) ports.UserRepository {
	return &UserRepositoryAdapter{
		DB: db,
	}
}

func (ur *UserRepositoryAdapter) CreateUser(userData *domain.User) (*domain.User, error) {
	query := `INSERT INTO users(id, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := ur.DB.Exec(query, userData.ID, userData.Name, userData.Email, userData.Password, userData.CreatedAt, userData.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new user: %v", err)
	}

	return userData, nil
}

func (ur *UserRepositoryAdapter) GetAllUsers() ([]domain.User, error) {
	var users []domain.User

	query := `SELECT id, name, email, password, created_at, updated_at FROM users`
	rows, err := ur.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	return users, nil
}

func (ur *UserRepositoryAdapter) GetUserByID(id string) (*domain.User, error) {
	var user domain.User

	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?`
	row := ur.DB.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}

		return nil, fmt.Errorf("failed to get user data based on id: %v", err)
	}

	return &user, nil
}

func (ur *UserRepositoryAdapter) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User

	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?`
	row := ur.DB.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user by email: %v", err)
	}

	return &user, nil
}

func (ur *UserRepositoryAdapter) UpdateUserByID(id string, userData *domain.User) (*domain.User, error) {
	var query string
	var args []interface{}

	if userData.Password == "" {
		query = `UPDATE users SET name = ?, email = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
		args = []interface{}{userData.Name, userData.Email, id}
	} else {
		query = `UPDATE users SET name = ?, email = ?, password = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
		args = []interface{}{userData.Name, userData.Email, userData.Password, id}
	}

	_, err := ur.DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update user data : %v", err)
	}

	return userData, nil
}

func (ur *UserRepositoryAdapter) DeleteUserByID(id string) error {
	query := `DELETE FROM users WHERE id = ?`
	result, err := ur.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check delete result: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
