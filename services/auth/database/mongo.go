package database

import (
	"context"
	"log"

	"github.com/lazbord/SpotyGo/services/auth/client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const USER_COLLECTION = "user"

type Adapter struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewAdapter(connectionURI string) (*Adapter, error) {
	dbName := "users"
	client, err := client.NewMongoClient(connectionURI)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &Adapter{
		client:   client,
		database: db,
	}, nil
}

func CreateUser(a *Adapter) {
	_, err := a.database.Collection(USER_COLLECTION).InsertOne(context.Background(), bson.M{"hello": "world"})
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}
}
