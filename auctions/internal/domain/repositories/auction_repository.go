package repositories

import "github.com/lioarce01/auction-microservices/internal/domain/entities"

type AuctionRepository interface {
	List() ([]entities.Auction, error)
	GetOne(id string) (entities.Auction, error)
	Create(auction entities.Auction) (entities.Auction, error)
	Update(creatorID string, auction entities.Auction) (entities.Auction, error)
	Delete(creatorID string, auctionID string) error
}
