package mongo

import (
	"context"
	"skynet/pkg"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Session struct {
	client *mongo.Client
}

func NewSession(config *root.MongoConfig) (*Session, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + config.Ip)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	return &Session{client}, nil
}

func (session *Session) Close() error {
	return session.client.Disconnect(context.TODO())
}

func (session *Session) DropDatabase(dbName string) error {
	db := session.client.Database(dbName)

	return db.Drop(context.TODO())
}
