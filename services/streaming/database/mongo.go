package database

import (
	"context"

	"github.com/lazbord/SpotyGo/common/client"
	"github.com/lazbord/SpotyGo/common/model"
	"go.mongodb.org/mongo-driver/bson"
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

func (a *Adapter) DBGetMusicByID(id string) (*model.Music, error) {
	collection := a.database.Collection(MUSIC_COLLECTION)
	music := model.Music{}
	err := collection.FindOne(context.Background(), bson.M{"videoid": id}).Decode(&music)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		return nil, err
	}

	return &music, nil
}
