package db_repositories

import (
	"context"
	"time"

	"github.com/lioarce01/auction-microservices/internal/domain/entities"
	"github.com/lioarce01/auction-microservices/internal/domain/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoAuctionRepository struct {
	DB *mongo.Database
}

func NewMongoAuctionRepository(db *mongo.Database) repositories.AuctionRepository {
	return &MongoAuctionRepository{DB: db}
}

func (r *MongoAuctionRepository) List() ([]entities.Auction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := r.DB.Collection("auctions")

	filter := bson.D{}

	findOptions := options.Find()

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var auctions []entities.Auction
	for cursor.Next(ctx) {
		var auction entities.Auction
		if err := cursor.Decode(&auction); err != nil {
			return nil, err
		}
		auctions = append(auctions, auction)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return auctions, nil
}

func (r *MongoAuctionRepository) GetOne(id string) (entities.Auction, error) {
	panic("unimplemented")
}

func (r *MongoAuctionRepository) Create(auction entities.Auction) (entities.Auction, error) {
	collection := r.DB.Collection("auctions")

	_, err := collection.InsertOne(context.TODO(), auction)
	if err != nil {
		return entities.Auction{}, err
	}

	return auction, nil
}

func (r *MongoAuctionRepository) Update(creatorID string, auction entities.Auction) (entities.Auction, error) {
	panic("unimplemented")
}

func (r *MongoAuctionRepository) Delete(creatorID string, auctionID string) error {
	panic("unimplemented")
}
