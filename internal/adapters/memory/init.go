package memory

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	db             *mongo.Database
	client         *mongo.Client
	taskCollection *mongo.Collection
}

func NewDatabase(ctx context.Context, connStr, dbName string) (*Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, fmt.Errorf("mongo client creation: %w", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("mongo client connection: %w", err)
	}
	db := client.Database(dbName)
	taskCollection := db.Collection("task")

	return &Database{
		db:             db,
		client:         client,
		taskCollection: taskCollection,
	}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.client.Ping(ctx, readpref.Primary())
}

func (d *Database) Close(ctx context.Context) error {
	return d.client.Disconnect(ctx)
}
