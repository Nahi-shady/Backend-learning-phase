package usecases

import (
	"task-manager/domain"
	"task-manager/repositories"

	"golang.org/x/crypto/bcrypt"
)

type JWTService interface {
	GenerateToken(userID string) (string, error)
	ParseToken(tokenString string) (string, error)
}

type UserUsecase struct {
	userRepo repositories.UserRepository
	jwtSvc   JWTService
}

func NewUserUsecase(userRepo repositories.UserRepository, jwtSvc JWTService) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
		jwtSvc:   jwtSvc,
	}
}

func (u *UserUsecase) Register(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &domain.User{
		Username: username,
		Password: string(hashedPassword),
	}
	return u.userRepo.CreateUser(user)
}

func (u *UserUsecase) LoginUser(username, password string) (string, error) {
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token, err := u.jwtSvc.GenerateToken(user.ID.Hex())
	if err != nil {
		return "", err
	}
	return token, nil
}
