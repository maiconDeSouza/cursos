package services

import (
	"errors"
	"meu-primeiro-crud-go-lang/internal/config"
	errapp "meu-primeiro-crud-go-lang/internal/config/err-app"
	"meu-primeiro-crud-go-lang/internal/models"
	"meu-primeiro-crud-go-lang/internal/repositories"
	"net/http"
	"time"
	"uuid"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
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

func (s *Services) generateJWT(id uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.cfg.Env.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *Services) generateCookie(jwt string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "session_token_user",
		Value:    jwt,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	return cookie
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

func (s *Services) LoginUser(userLogin models.UserLogin) (*models.UserResponse, *http.Cookie, *errapp.ErrResponse) {
	if err := s.validate.Struct(userLogin); err != nil {
		errApp := errapp.NewBadRequestError("Problema nos campos", err)
		return nil, nil, errApp
	}

	user, err := s.repo.FindByEmail(userLogin.Email)
	if err != nil {
		if errors.Is(err, errapp.ErrUserNotFound) {
			errApp := errapp.NewBadRequestError("Senha ou Login inválidos", err)
			return nil, nil, errApp
		}
		errApp := errapp.NewInternalServerError("Erro do servidor", err)
		return nil, nil, errApp
	}

	ok := s.checkPasswordHash(userLogin.Password, user.Password)
	if !ok {
		errApp := errapp.NewBadRequestError("Senha ou Login inválidos", err)
		return nil, nil, errApp
	}

	jtw, err := s.generateJWT(user.ID)
	if err != nil {
		errApp := errapp.NewInternalServerError("Erro do servidor", err)
		return nil, nil, errApp
	}

	cookie := s.generateCookie(jtw)
	userResponse := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	}

	return &userResponse, cookie, nil
}

func (s *Services) UpdateUser(id uuid.UUID, userUpdate models.UserUpdate) (*models.UserResponse, *errapp.ErrResponse) {
	user := models.User{}
	if err := s.validate.Struct(userUpdate); err != nil {
		errApp := errapp.NewBadRequestError("Problema nos campos", err)
		return nil, errApp
	}

	if userUpdate.Password != nil {
		if userUpdate.Password != userUpdate.RepeatPassword {
			errApp := errapp.NewBadRequestError("A senhas devem ser iguais", nil)
			return nil, errApp
		}
		hash, err := s.hashPassword(*userUpdate.Password)
		if err != nil {
			errApp := errapp.NewInternalServerError("Erro do servidor", err)
			return nil, errApp
		}
		user.Password = hash
	}

	if userUpdate.Name != nil {
		user.Name = *userUpdate.Name
	}

	if userUpdate.Age != nil {
		user.Age = *userUpdate.Age
	}

	if userUpdate.Email != nil {
		user.Email = *userUpdate.Email
	}

	err := s.repo.Update(id, user)
	if err != nil {
		if errors.Is(err, errapp.ErrUserNotFound) {
			errApp := errapp.NewBadRequestError("Senha ou Login inválidos", err)
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

func (s *Services) DeleteUser(id uuid.UUID) *errapp.ErrResponse {
	err := s.repo.Delete(id)
	if err != nil {
		if errors.Is(err, errapp.ErrUserNotFound) {
			errApp := errapp.NewBadRequestError("Senha ou Login inválidos", err)
			return errApp
		}
		errApp := errapp.NewInternalServerError("Erro do servidor", err)
		return errApp
	}

	return nil
}
