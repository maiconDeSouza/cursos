package model

import "uuid"

type UserApp struct {
	ID       uuid.UUID `bson:"_id"`
	Name     string    `bson:"name"`
	Age      uint8     `bson:"age"`
	Email    string    `bson:"email"`
	Password string    `bson:"password"`
	Admin    bool      `bson:"admin"`
}

type Request struct {
	Name           string `json:"name" validate:"required,min=4,max=25"`
	Email          string `json:"email" validate:"required,email"`
	Age            uint8  `json:"age" validate:"required,min=1,max=140"`
	Password       string `json:"password" validate:"required,min=4,max=55"`
	RepeatPassword string `json:"repeatPassword" validate:"required,min=4,max=55"`
}

type Response struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Age   uint8     `json:"age"`
}
