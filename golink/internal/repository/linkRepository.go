package repository

import (
	"context"
	"errors"
	"golink/internal/config"
	"golink/internal/database"
	errapp "golink/internal/errApp"
	"golink/internal/model"
	"time"
	"uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LinkRepository struct {
	collection *mongo.Collection
	env        *config.ENV
}

func NewLinkRepository(collection *database.Collection, env *config.ENV) *LinkRepository {
	return &LinkRepository{
		collection: collection.Links,
		env:        env,
	}
}

func (r *LinkRepository) Create(link model.Link) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	if _, err := r.collection.InsertOne(ctx, link); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errapp.ErrTitleAlreadyInLink
		}
		return err
	}

	return nil
}

func (r *LinkRepository) FindByTitle(title string, link *model.Link) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	filter := bson.M{"title": title}
	if err := r.collection.FindOne(ctx, filter).Decode(link); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errapp.ErrLinkNotFound
		}
		return err
	}
	return nil
}

func (r *LinkRepository) FindByID(id uuid.UUID, link *model.Link) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	if err := r.collection.FindOne(ctx, filter).Decode(link); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errapp.ErrLinkNotFound
		}
		return err
	}
	return nil
}

func (r *LinkRepository) CountByUserID(user_id uuid.UUID) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	filter := bson.M{"user_id": user_id}
	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *LinkRepository) Update(id uuid.UUID, link *model.Link) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": link}

	re, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if re.MatchedCount == 0 {
		return errapp.ErrLinkNotFound
	}

	return nil
}

func (r *LinkRepository) Delete(id uuid.UUID) error {
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
