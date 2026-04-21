package services

import (
	"database/sql"
	// "drinklytics/internal/helper"
	// "drinklytics/internal/middleware"
	// "drinklytics/internal/models"
	"fmt"
	"log"
	"time"
)

type userService struct {
	db *sql.DB
}

type UserService interface {
	CreateUser(user string) (string, error)
	// UpdateUser(request *models.UserRequest) (string, error)
	// UpdateUserPassword(request *models.UserRequest) (string, error)
	// DeleteUser(request *models.UserRequest) error

	// GetUser(userId int64) ([]models.UserDetails, error)
	// GetAllUsers() ([]models.UserDetails, error)
	// HardDeleteUser(request *models.UserRequest) error
	// UpdateUserRole(request *models.User) (string, error)
	// UpdateUserStatus(request *models.User) (string, error)
}

func NewUserService(db *sql.DB) UserService {
	return &userService{db}
}

func (s *userService) CreateUser(user string) (string, error) {
	var result bool
	query := `SELECT
    			CASE WHEN EXISTS
       				(
           				SELECT * FROM user_data WHERE name = $1 AND is_active = true
               		)
                THEN 'TRUE'
                ELSE 'FALSE'
            END`
    err := s.db.QueryRow(query,user).Scan(&result)

	if err != nil {
		return "User Creation Failed", err
	}else if result{
		return "User already exists", fmt.Errorf("User already exists")
	}

	query = `INSERT INTO user_data (name, role, is_active, created_at)
		VALUES ($1, $2, $3, $4)`

	_, err = s.db.Exec(query, user, "viewer", true, time.Now())

	if err != nil {
		log.Println("Error Executing query: ", err)
		return "User Creation Failed", err
	}
	return user, nil
}

// func (s *userService) CreateUser(request *models.UserRequest) (string, error) {

// 	hashPassword, err := middleware.HashPassword(request.Password)

// 	if err != nil {
// 		return "User Creation Failed", err
// 	}

// 	query := `INSERT INTO users (name, email, password, role, is_active, created_at)
// 		VALUES ($1, $2, $3, $4, $5, $6)`

// 	_, err = s.db.Exec(query, request.Name, request.Email, hashPassword, "viewer", true, time.Now())

// 	if err != nil {
// 		log.Println("Error Executing query: ", err)
// 		return "User Creation Failed", err
// 	}
// 	return "Create Coupons Successful", nil
// }

// func (s *userService) UpdateUser(request *models.UserRequest) (string, error) {
// 	query := `UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4 AND (name, email) IS DISTINCT FROM ($1,$2)`

// 	_, err := s.db.Exec(query, request.Name, request.Email, time.Now(), request.ID)
// 	if err != nil {
// 		log.Println("Error Executing query:", err)
// 		return "User Upation Failed", err
// 	}

// 	return "Successfull updated", nil
// }

// func (s *userService) UpdateUserPassword(request *models.UserRequest) (string, error) {
// 	hashPassword, err := middleware.HashPassword(request.Password)

// 	if err != nil {
// 		return "User Password Upation Failed", err
// 	}

// 	query := `UPDATE users SET password = $1,updated_at = $2 WHERE id = $3`

// 	_, err = s.db.Exec(query, hashPassword, time.Now(), request.ID)
// 	if err != nil {
// 		log.Println("Error Executing query:", err)
// 		return "User Password Upation Failed", err
// 	}

// 	return "Successfull password updated", nil
// }

// func (s *userService) DeleteUser(request *models.UserRequest) error {
// 	query := `UPDATE users SET is_active = false WHERE id = $1`
// 	_, err := s.db.Exec(query, request.ID)
// 	if err != nil {
// 		log.Println("Error Executing query:", err)
// 	}
// 	return err
// }

// func (s *userService) UpdateUserRole(request *models.User) (string, error) {
// 	query := `UPDATE users SET role = $1, updated_at = $2 WHERE id = $3`

// 	_, err := s.db.Exec(query, request.Role, time.Now(), request.ID)
// 	if err != nil {
// 		log.Println("Error Executing query:", err)
// 		return "User Role Upation Failed", err
// 	}

// 	return "Successfull Role updated", nil
// }

// func (s *userService) UpdateUserStatus(request *models.User) (string, error) {
// 	query := `UPDATE users SET is_active = $1, updated_at = $2 WHERE id = $3`

// 	_, err := s.db.Exec(query, request.IsActive, time.Now(), request.ID)
// 	if err != nil {
// 		log.Println("Error Executing query:", err)
// 		return "User Role Upation Failed", err
// 	}

// 	return "Successfull status updated", nil
// }

// func (s *userService) GetUser(userId int64) ([]models.UserDetails, error) {
// 	var users []models.UserDetails
// 	query := `SELECT id, name, email, role, is_active, created_at, updated_at FROM users WHERE id = $1` // is_active = true
// 	rows, err := s.db.Query(query, userId)
// 	if err != nil {
// 		log.Println("Error Executing query:", err)
// 		return []models.UserDetails{}, err
// 	}

// 	err = helper.UsersDetailsDtoMapper(rows, &users)
// 	if err != nil {
// 		log.Println("Error Executing query:", err)
// 		return []models.UserDetails{}, err
// 	}
// 	return users, nil
// }

// func (s *userService) GetAllUsers() ([]models.UserDetails, error) {
// 	var users []models.UserDetails
// 	query := `SELECT id, name, email, role, is_active, created_at, updated_at FROM users` // is_active = true
// 	rows, err := s.db.Query(query)
// 	if err != nil {
// 		log.Println("Error Executing query:", err)
// 		return []models.UserDetails{}, fmt.Errorf("Error Executing query: %w", err)
// 	}
// 	err = helper.UsersDetailsDtoMapper(rows, &users)
// 	if err != nil {
// 		log.Println("Error Executing query:", err)
// 		return []models.UserDetails{}, err
// 	}
// 	return users, nil
// }

// func (s *userService) HardDeleteUser(request *models.UserRequest) error {
// 	query := `DELETE from users WHERE id = $1`
// 	_, err := s.db.Exec(query, request.ID)
// 	if err != nil {
// 		log.Println("Error Executing query:", err)
// 	}
// 	return err
// }
