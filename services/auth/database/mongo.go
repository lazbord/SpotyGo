package main

import (
	"github.com/lazbord/SpotyGo/services/auth/client"
	"go.mongodb.org/mongo-driver/mongo"
)

type Adapter struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewAdaptater(connectionURI string) (*Adapter, error) {
	dbName := "user"
	client, err := client.MongoClient(connectionURI)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &Adapter{
		client:   client,
		database: db,
	}, nil
}
