package service

import (
	errapp "golink/internal/errApp"
	"golink/internal/model"
	"time"
	"uuid"

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
