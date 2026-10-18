package repository

import (
	"context"
	"errors"
	"golink/internal/config"
	"golink/internal/database"
	errapp "golink/internal/errApp"
	"golink/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
	env        *config.ENV
}

func NewUserRepository(collection *database.Collection, env *config.ENV) *UserRepository {
	return &UserRepository{
		collection: collection.Users,
		env:        env,
	}
}

func (r *UserRepository) Create(user model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	if _, err := r.collection.InsertOne(ctx, user); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errapp.ErrEmailAlreadyInUse
		}
		return err
	}

	return nil
}

func (r *UserRepository) FindByEmail(email string, user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	filter := bson.M{"email": email}
	if err := r.collection.FindOne(ctx, filter).Decode(user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errapp.ErrUserNotFound
		}
		return err
	}

	return nil
}
