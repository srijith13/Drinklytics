package services

import (
	"database/sql"
	"drinklytics/internal/config"
	"drinklytics/internal/helper"
	"drinklytics/internal/middleware"
	"drinklytics/internal/models"
	"fmt"
	"log"
)

type authService struct {
	db *sql.DB
}

type AuthService interface {
	LoginUser(request *models.UserRequest) (string, error)
	LogoutUser(userId int64) error
}

func NewAuthService(db *sql.DB) AuthService {
	return &authService{db}
}

func (s *authService) LoginUser(request *models.UserRequest) (string, error) {
	var user models.User
	query := `SELECT * FROM users WHERE email = $1 AND is_active = true `
	row := s.db.QueryRow(query, request.Email)

	helper.UserDtoMapper(row, &user)

	checkPassword := middleware.CheckPasswordHash(request.Password, user.Password)

	if !checkPassword {
		err := fmt.Errorf("Wrong Password")
		log.Println("Login Failed: ", err)
		return "Login Failed", err
	}

	token, err := middleware.TokenGenerator(&user)

	if err != nil {
		log.Println("Login Failed: ", err)
		return "Login Failed", err
	}

	config.UserCache[user.ID] = token

	return token, nil

}

func (s *authService) LogoutUser(userId int64) error {
	if _, ok := config.UserCache[userId]; !ok {
		return fmt.Errorf("Missing token")
	}
	fmt.Println("before cahce ", config.UserCache[userId])
	delete(config.UserCache, userId)
	fmt.Println("after cahce ", config.UserCache[userId])

	return nil
}
