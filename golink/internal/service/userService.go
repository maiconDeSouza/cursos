package service

import (
	"errors"
	"golink/internal/config"
	errapp "golink/internal/errApp"
	"golink/internal/model"
	"golink/internal/repository"
	"net/http"

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

func (s *UserService) Login(login *model.Login) (*http.Cookie, *model.LoginResponse, *errapp.ErrResponse) {
	if err := s.validate.Struct(login); err != nil {
		return nil, nil, errapp.NewBadRequestError("Problema em um dos campos: "+err.Error(), err)
	}

	user := model.User{}

	if err := s.repo.FindByEmail(login.Email, &user); err != nil {
		if errors.Is(err, errapp.ErrUserNotFound) {
			return nil, nil, errapp.NewBadRequestError("Usuário ou senha inválidos", nil)
		}
		return nil, nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	if ok := checkPasswordHash(login.Password, user.Password); !ok {
		return nil, nil, errapp.NewBadRequestError("Usuário ou senha inválidos", nil)
	}

	jwtString, err := generateJWT(user.ID, s.env)
	if err != nil {
		return nil, nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	cookie := generateCookie(jwtString)
	loginResponse := model.LoginResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	return cookie, &loginResponse, nil
}
