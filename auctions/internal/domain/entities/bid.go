package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Bid struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	AuctionID bson.ObjectID `bson:"auction_id" json:"auction_id"`
	UserID    string        `bson:"user_id" json:"user_id"`
	Price     float64       `bson:"price" json:"price"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
}
