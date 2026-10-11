package repositories

import (
	"context"
	"meu-primeiro-crud-go-lang/internal/config"
	errapp "meu-primeiro-crud-go-lang/internal/config/err-app"
	"meu-primeiro-crud-go-lang/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoriesInterface interface {
	Create(user *models.User) error
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	filter := bson.M{"email": user.Email}
	result := r.collection.FindOne(ctx, filter)
	if result.Err() == mongo.ErrNoDocuments {
		_, err := r.collection.InsertOne(ctx, user)
		return err
	}
	return errapp.ErrEmailAlreadyInUse
}
