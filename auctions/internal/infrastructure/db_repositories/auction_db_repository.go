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

	auctions := make([]entities.Auction, 0)
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return entities.Auction{}, err
	}

	collection := r.DB.Collection("auctions")

	filter := bson.M{"_id": objectId}

	var auction entities.Auction

	err = collection.FindOne(ctx, filter).Decode(&auction)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entities.Auction{}, nil
		}
		return entities.Auction{}, err
	}

	return auction, nil
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := r.DB.Collection("auctions")

	objectId, err := bson.ObjectIDFromHex(auction.ID.Hex())
	if err != nil {
		return entities.Auction{}, err
	}

	filter := bson.M{"_id": objectId, "creator_id": creatorID}

	update := bson.M{
		"$set": bson.M{
			"title":         auction.Title,
			"description":   auction.Description,
			"current_price": auction.CurrentPrice,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return entities.Auction{}, err
	}

	if result.MatchedCount == 0 {
		return entities.Auction{}, mongo.ErrNoDocuments
	}

	return auction, nil
}

func (r *MongoAuctionRepository) Delete(creatorID string, auctionID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := r.DB.Collection("auctions")

	objectId, err := bson.ObjectIDFromHex(auctionID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId, "creator_id": creatorID}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
