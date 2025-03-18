package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Auction struct {
	ID           bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title        string        `bson:"title" json:"title"`
	Description  string        `bson:"description,omitempty" json:"description,omitempty"`
	CreatorID    bson.ObjectID `bson:"creator_id" json:"creator_id"`
	CurrentPrice float64       `bson:"current_price" json:"current_price"`
	CreatedAt    time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time     `bson:"updated_at" json:"updated_at"`
	EndDate      time.Time     `bson:"end_date,omitempty" json:"end_date,omitempty"`
	Status       string        `bson:"status,omitempty" json:"status,omitempty"`
}

type AuctionRequest struct {
	Title        string    `bson:"title" json:"title"`
	Description  string    `bson:"description,omitempty" json:"description,omitempty"`
	CurrentPrice float64   `bson:"current_price" json:"current_price"`
	EndDate      time.Time `bson:"end_date,omitempty" json:"end_date,omitempty"`
	Status       string    `bson:"status,omitempty" json:"status,omitempty"`
}
