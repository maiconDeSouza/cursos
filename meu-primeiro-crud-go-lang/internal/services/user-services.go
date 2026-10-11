package services

import (
	"errors"
	"meu-primeiro-crud-go-lang/internal/config"
	errapp "meu-primeiro-crud-go-lang/internal/config/err-app"
	"meu-primeiro-crud-go-lang/internal/models"
	"meu-primeiro-crud-go-lang/internal/repositories"
	"uuid"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type ServicesInterface interface {
	CreateUser(UserRequest *models.UserRequest) (*models.UserResponse, *errapp.ErrResponse)
}

type Services struct {
	repo     repositories.RepositoriesInterface
	cfg      *config.Config
	validate *validator.Validate
}

func NewServices(repo repositories.RepositoriesInterface, cfg *config.Config) *Services {
	return &Services{repo: repo, cfg: cfg, validate: validator.New()}
}

func (s *Services) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (s *Services) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func (s *Services) CreateUser(UserRequest *models.UserRequest) (*models.UserResponse, *errapp.ErrResponse) {
	if err := s.validate.Struct(UserRequest); err != nil {
		errApp := errapp.NewBadRequestError("Problema nos campos", err)
		return nil, errApp
	}

	if UserRequest.Password != UserRequest.RepeatPassword {
		errApp := errapp.NewBadRequestError("As senhas são diferentes", nil)
		return nil, errApp
	}

	h, err := s.hashPassword(UserRequest.Password)
	if err != nil {
		errApp := errapp.NewInternalServerError("Erro do servidor", err)
		return nil, errApp
	}

	user := models.User{
		ID:       uuid.New(),
		Name:     UserRequest.Name,
		Email:    UserRequest.Email,
		Age:      UserRequest.Age,
		Password: h,
	}

	err = s.repo.Create(&user)
	if err != nil {
		if errors.Is(err, errapp.ErrEmailAlreadyInUse) {
			errApp := errapp.NewBadRequestError("Email já existe", err)
			return nil, errApp
		}
		errApp := errapp.NewInternalServerError("Erro do servidor", err)
		return nil, errApp
	}

	userResponse := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	}

	return &userResponse, nil
}
