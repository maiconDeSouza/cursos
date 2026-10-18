package database

import (
	"context"
	"fmt"
	"golink/internal/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	Users *mongo.Collection
	Links *mongo.Collection
}

func NewConnectMongo(env *config.ENV) *Collection {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@localhost:%s/",
		env.MONGO_ROOT_USER,
		env.MONGO_ROOT_PASSWORD,
		env.MONGO_PORT,
	)
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌​ Erro ao ligar ao MongoDB.", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("❌​ Erro ao ligar ao MongoDB.", err)
	}

	fmt.Println("✅ Conectado ao banco de dados com sucesso!")

	collection := client.Database(env.MONGO_DB_NAME)
	usersCollection := collection.Collection(env.MONGO_COLLECTION_USERS)
	linksCollection := collection.Collection(env.MONGO_COLLECTION_LINKS)

	userIndexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
	if _, err := usersCollection.Indexes().CreateMany(ctx, userIndexes); err != nil {
		log.Fatal("❌ Erro ao criar índice único para e-mail: ", err)
	}

	linkIndexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "title", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
	if _, err := linksCollection.Indexes().CreateMany(ctx, linkIndexes); err != nil {
		log.Fatal("❌ Erro ao criar índice único para o título do link: ", err)
	}

	c := &Collection{
		Users: usersCollection,
		Links: linksCollection,
	}

	return c
}
