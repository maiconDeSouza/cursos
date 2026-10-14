package service

import (
	"errors"
	"meu-primeiro-crud-go-lang-v2/internal/config"
	"meu-primeiro-crud-go-lang-v2/internal/errapp"
	"meu-primeiro-crud-go-lang-v2/internal/model"
	"meu-primeiro-crud-go-lang-v2/internal/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	repo     repository.RepoInterface
	env      *config.Env
	errapp   *errapp.ErrApp
	validate *validator.Validate
}

func NewAdminService(repo repository.RepoInterface, env *config.Env, errapp *errapp.ErrApp) *AdminService {
	return &AdminService{repo: repo, env: env, errapp: errapp, validate: validator.New()}
}

func (s *AdminService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (s *AdminService) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func (s *AdminService) NewAdmin(adminRequest model.Request) (*model.Response, *errapp.ErrAppResult) {
	if err := s.validate.Struct(adminRequest); err != nil {
		return nil, s.errapp.BadRequestError("Problema em um dos campos", err)
	}

	if adminRequest.Password != adminRequest.RepeatPassword {
		return nil, s.errapp.BadRequestError("As senhas tem que ser iguais", nil)
	}

	h, err := s.hashPassword(adminRequest.Password)
	if err != nil {
		return nil, s.errapp.InternalServerError("Problema no servidor", err)
	}

	user := model.UserApp{
		Name:     adminRequest.Name,
		Age:      adminRequest.Age,
		Email:    adminRequest.Email,
		Password: h,
		Admin:    true,
	}

	err = s.repo.Create(&user)
	if err != nil {
		if errors.Is(err, s.errapp.ErrEmailAlreadyInUse) {
			return nil, s.errapp.BadRequestError("Email em uso", err)
		}
		return nil, s.errapp.InternalServerError("Problema no servidor", err)
	}

	res := model.Response{
		ID:    user.ID,
		Name:  user.Name,
		Age:   user.Age,
		Email: user.Email,
	}

	return &res, nil
}
