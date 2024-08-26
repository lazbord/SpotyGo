package database

import (
	"context"

	"github.com/lazbord/SpotyGo/common/client"
	"github.com/lazbord/SpotyGo/common/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const AUTHCOLLECTION = "user"

type Adapter struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewAdapter(connectionURI string) (*Adapter, error) {
	dbName := "Auth"
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

func (a *Adapter) CreateAuth(auth model.Auth) (string, error) {

	auth.UserPassword, _ = HashPassword(auth.UserPassword)

	collection := a.database.Collection(AUTHCOLLECTION)
	res, err := collection.InsertOne(context.Background(), auth, nil)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (a Adapter) GetAuthByEmail(email string) (*model.Auth, error) {
	collection := a.database.Collection(AUTHCOLLECTION)
	auth := model.Auth{}
	err := collection.FindOne(context.Background(), bson.M{"user_email": email}).Decode(&auth)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		return nil, err
	}

	return &auth, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
