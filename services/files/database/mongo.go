package database

import (
	"context"
	"errors"

	"github.com/lazbord/SpotyGo/common/client"
	"github.com/lazbord/SpotyGo/common/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (a *Adapter) DBAddMusic(music model.Music) (string, error) {
	collection := a.database.Collection(MUSIC_COLLECTION)
	res, err := collection.InsertOne(context.Background(), music, nil)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (a *Adapter) DBDeleteMusic(id string) error {
	collection := a.database.Collection(MUSIC_COLLECTION)

	result, err := collection.DeleteOne(context.Background(), bson.M{"videoid": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("mongo : no document found")
	}

	return nil
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
