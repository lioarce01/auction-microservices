package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Auction struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Title        string        `bson:"title"`
	Description  string        `bson:"description,omitempty"`
	CreatorID    bson.ObjectID `bson:"creator_id"`
	CurrentPrice float64       `bson:"current_price"`
	CreatedAt    time.Time     `bson:"created_at"`
	UpdatedAt    time.Time     `bson:"updated_at"`
	EndDate      time.Time     `bson:"end_date,omitempty"`
	Status       string        `bson:"status,omitempty"`
}
