package repository

import (
	"context"
	"meu-primeiro-crud-go-lang-v2/internal/errapp"
	"meu-primeiro-crud-go-lang-v2/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db     *mongo.Collection
	errapp *errapp.ErrApp
}

func NewRepo(db *mongo.Collection, errapp *errapp.ErrApp) *Repository {
	return &Repository{db: db, errapp: errapp}
}

func (r *Repository) Create(user *model.UserApp) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	_, err := r.db.InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return r.errapp.ErrEmailAlreadyInUse
		}
		return err
	}

	return nil
}
