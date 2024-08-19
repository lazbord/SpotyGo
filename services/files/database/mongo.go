package database

import (
	"context"

	"github.com/lazbord/SpotyGo/model"
	"github.com/lazbord/SpotyGo/services/files/client"
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

func (a *Adapter) DBRemoveMusic(music model.Music) error {
	collection := a.database.Collection(MUSIC_COLLECTION)
	_, err := collection.DeleteOne(context.Background(), music, nil)
	if err != nil {
		return err
	}
	return nil
}

func (a *Adapter) DBGetMusicByID(id string) (*model.Music, error) {
	collection := a.database.Collection(MUSIC_COLLECTION)
	music := model.Music{}
	err := collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&music)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		return nil, err
	}

	return &music, nil
}
