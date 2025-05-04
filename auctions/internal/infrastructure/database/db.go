package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewDBConnection() (*mongo.Database, error) {
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No se encontró el archivo .env, se utilizarán las variables de entorno definidas")
		}
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set")
	}

	dbName := os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		log.Fatal("MONGO_DB_NAME is not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI).
		SetConnectTimeout(10 * time.Second).
		SetMaxPoolSize(10)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error pinging MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB")

	return client.Database(dbName), nil
}
