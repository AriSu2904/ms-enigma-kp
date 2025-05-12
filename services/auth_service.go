package services

import (
	"awesomeProject/models"
	"awesomeProject/repositories"
	"awesomeProject/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(credentials *models.LoginRequest) (*models.AuthResponse, error)
}

type authService struct {
	userRepository repositories.UserRepository
	config         *utils.JWTConfig
}

func NewAuthService(userRepo repositories.UserRepository, config *utils.JWTConfig) AuthService {
	return &authService{
		userRepository: userRepo,
		config:         config,
	}
}

func (s *authService) Login(req *models.LoginRequest) (*models.AuthResponse, error) {
	user, err := s.userRepository.FindByNik(req.Nik)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate token
	token, err := s.config.GenerateToken(user.Nik)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}
