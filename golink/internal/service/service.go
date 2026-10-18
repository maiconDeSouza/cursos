package service

import (
	"golink/internal/config"
	errapp "golink/internal/errApp"
	"golink/internal/model"
	"net/http"
	"time"
	"uuid"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Register(userRequest *model.UserRequest) (*model.UserResponse, *errapp.ErrResponse)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func generateJWT(id uuid.UUID, env *config.ENV) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(env.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateCookie(jwt string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "session_token_user",
		Value:    jwt,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	return cookie
}

func userRequestToUser(userRequest *model.UserRequest) *model.User {
	user := model.User{
		ID:        uuid.New(),
		Email:     userRequest.Email,
		Password:  userRequest.Password,
		CreatedAt: time.Now(),
	}
	return &user
}

func userToUserResponse(user *model.User) *model.UserResponse {
	userResponse := model.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
	return &userResponse
}
