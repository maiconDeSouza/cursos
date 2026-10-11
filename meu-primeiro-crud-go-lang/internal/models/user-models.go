package models

import "uuid"

type User struct {
	ID       uuid.UUID `bson:"_id"`
	Name     string    `bson:"name"`
	Email    string    `json:"email"`
	Age      uint8     `json:"age"`
	Password string    `json:"password"`
}

type UserRequest struct {
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Age            uint8  `json:"age" validate:"required,min=1,max=140"`
	Password       string `json:"password" validate:"required,min=4,max=55"`
	RepeatPassword string `json:"repeatPassword" validate:"required,min=4,max=55"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Age   uint8     `json:"age"`
}
