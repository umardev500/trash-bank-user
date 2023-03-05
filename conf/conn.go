package conf

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewConn creates a new connection to a MongoDB database
func NewConn() *mongo.Database {
	// Set context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connStr := os.Getenv("CONN_STR")

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}

	// Ping to the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("db_trash_bank")

	return db
}
