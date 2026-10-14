package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	db *mongo.Collection
}

func NewRepo(db *mongo.Collection) *Repository {
	return &Repository{db: db}
}
