package service

import (
	"errors"
	"golink/internal/config"
	errapp "golink/internal/errApp"
	"golink/internal/model"
	"golink/internal/repository"

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	repo     repository.UserRepositoryInterface
	env      *config.ENV
	validate *validator.Validate
}

func NewUserService(repo repository.UserRepositoryInterface, env *config.ENV) *UserService {
	return &UserService{repo: repo, env: env, validate: validator.New()}
}

func (s *UserService) Register(userRequest *model.UserRequest) (*model.UserResponse, *errapp.ErrResponse) {
	if err := s.validate.Struct(userRequest); err != nil {
		return nil, errapp.NewBadRequestError("Problema em um dos campos: "+err.Error(), err)
	}

	if userRequest.Password != userRequest.RepeatPassword {
		return nil, errapp.NewBadRequestError("As senhas tem que ser iguais", nil)
	}

	h, err := hashPassword(userRequest.Password)
	if err != nil {
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}
	userRequest.Password = h

	user := userRequestToUser(userRequest)

	if err := s.repo.Create(*user); err != nil {
		if errors.Is(err, errapp.ErrEmailAlreadyInUse) {
			return nil, errapp.NewBadRequestError("Email em uso: "+err.Error(), err)
		}
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	userResponse := userToUserResponse(user)

	return userResponse, nil
}
