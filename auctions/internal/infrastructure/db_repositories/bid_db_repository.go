package db_repositories

import (
	"context"
	"errors"
	"time"

	"github.com/lioarce01/auction-microservices/internal/domain/entities"
	"github.com/lioarce01/auction-microservices/internal/domain/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoBidRepository struct {
	DB *mongo.Database
}

func NewMongoBidRepository(db *mongo.Database) repositories.BidRepository {
	return &MongoBidRepository{DB: db}
}

func (r *MongoBidRepository) CreateBid(auctionID string, userID string, price float64) (entities.Bid, error) {
	auctionCollection := r.DB.Collection("auctions")
	bidsCollection := r.DB.Collection("bids")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	auctionObjID, err := bson.ObjectIDFromHex(auctionID)
	if err != nil {
		return entities.Bid{}, errors.New("invalid auction ID format")
	}

	var auction entities.Auction
	err = auctionCollection.FindOne(ctx, bson.M{"_id": auctionObjID}).Decode(&auction)
	if err != nil {
		return entities.Bid{}, err
	}

	if price <= auction.CurrentPrice {
		return entities.Bid{}, errors.New("el precio de la puja debe ser mayor al precio actual de la subasta")
	}

	bid := entities.Bid{
		ID:        bson.NewObjectID(),
		AuctionID: auctionObjID,
		UserID:    userID,
		Price:     price,
		CreatedAt: time.Now(),
	}

	_, err = bidsCollection.InsertOne(ctx, bid)
	if err != nil {
		return entities.Bid{}, err
	}

	update := bson.M{
		"$set": bson.M{
			"current_price": price,
			"updated_at":    time.Now(),
		},
	}
	_, err = auctionCollection.UpdateOne(ctx, bson.M{"_id": auctionObjID, "current_price": bson.M{"$lt": price}}, update)
	if err != nil {
		return entities.Bid{}, err
	}

	return bid, nil
}

func (r *MongoBidRepository) GetBids(AuctionID string) ([]entities.Bid, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := r.DB.Collection("bids")

	filter := bson.D{}

	findOptions := options.Find()

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	bids := make([]entities.Bid, 0)
	for cursor.Next(ctx) {
		var bid entities.Bid
		if err := cursor.Decode(&bid); err != nil {
			return nil, err
		}
		bids = append(bids, bid)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return bids, nil
}
