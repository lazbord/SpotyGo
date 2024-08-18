package database

import (
	"github.com/lazbord/SpotyGo/services/streaming/client"
	"go.mongodb.org/mongo-driver/mongo"
)

const MUSIC_COLLECTION = "music"

type Adapter struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewAdapter(connectionURI string) (*Adapter, error) {
	dbName := "Streaming"
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
