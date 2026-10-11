package repositories

import (
	"context"
	"errors"
	"meu-primeiro-crud-go-lang/internal/config"
	errapp "meu-primeiro-crud-go-lang/internal/config/err-app"
	"meu-primeiro-crud-go-lang/internal/models"
	"time"
	"uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoriesInterface interface {
	Create(user *models.User) error
	FindByID(id uuid.UUID) (*models.User, error)
}

type Repositories struct {
	client     *mongo.Client
	collection *mongo.Collection
	cfg        *config.Config
}

func NewRepositories(client *mongo.Client, cfg *config.Config) *Repositories {
	return &Repositories{
		client:     client,
		collection: client.Database(cfg.Env.MONGO_DB_NAME).Collection(cfg.Env.MONGO_COLLECTION),
		cfg:        cfg,
	}
}

func (r *Repositories) Create(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"email": user.Email}

	var existingUser struct{}
	err := r.collection.FindOne(ctx, filter).Decode(&existingUser)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			_, insertErr := r.collection.InsertOne(ctx, user)
			return insertErr
		}
		return err
	}

	return errapp.ErrEmailAlreadyInUse
}

func (r *Repositories) FindByID(id uuid.UUID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	filter := bson.M{"_id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errapp.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *Repositories) FindByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	filter := bson.M{"email": email}

	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errapp.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *Repositories) Update(id uuid.UUID, user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errapp.ErrUserNotFound
	}

	return nil
}

func (r *Repositories) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errapp.ErrUserNotFound
	}

	return nil
}
