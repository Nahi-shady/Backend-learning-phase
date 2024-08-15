package services

import (
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/repositories"
	"task-manager/internal/infrastructure/auth"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repositories.UserRepository
	jwtSvc   auth.JWTService
}

func NewUserService(userRepo repositories.UserRepository, jwtSvc auth.JWTService) *UserService {
	return &UserService{
		userRepo: userRepo,
		jwtSvc:   jwtSvc,
	}
}

func (s *UserService) Register(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &entities.User{
		Username: username,
		Password: string(hashedPassword),
	}
	return s.userRepo.CreateUser(user)
}

func (s *UserService) LoginUser(username, password string) (string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	return s.jwtSvc.GenerateToken(user.ID.Hex())
}
