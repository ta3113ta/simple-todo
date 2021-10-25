package todoCRUD

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout = 5
	connectionURI  = "mongodb://localhost:2717/projects"
)

func GetConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))

	if err != nil {
		log.Fatalf("Failed to create client : %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	if err = client.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect to cluster: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}

	return client, ctx, cancel
}