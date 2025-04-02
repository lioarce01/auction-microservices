package repositories

import "github.com/lioarce01/auction-microservices/internal/domain/entities"

type BidRepository interface {
	CreateBid(AuctionID string, UserID string, Price float64) (entities.Bid, error)
	GetBids(AuctionID string) ([]entities.Bid, error)
}
