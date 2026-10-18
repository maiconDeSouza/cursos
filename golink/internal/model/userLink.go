package model

import (
	"time"
	"uuid"
)

type User struct {
	ID        uuid.UUID `bson:"_id"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
}

type UserRequest struct {
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=5,max=25"`
	RepeatPassword string `json:"repeat_password" validate:"required,min=5,max=25"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Link struct {
	ID        uuid.UUID `bson:"_id"`
	Title     string    `bson:"title"`
	URL       string    `bson:"url"`
	UserID    uuid.UUID `bson:"user_id"`
	CreatedAt time.Time `bson:"created_at"`
}

type LinkReqAndRes struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title" validate:"required,min=3,max=50"`
	URL       string    `json:"url" validate:"required,url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
