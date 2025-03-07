package auction

import (
	"github.com/lioarce01/auction-microservices/internal/domain/entities"
	"github.com/lioarce01/auction-microservices/internal/domain/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateAuctionUseCase struct {
	AuctionRepo repositories.AuctionRepository
}

func NewCreateAuctionUseCase(repo repositories.AuctionRepository) *CreateAuctionUseCase {
	return &CreateAuctionUseCase{AuctionRepo: repo}
}

func (uc *CreateAuctionUseCase) Execute(auction entities.Auction) (entities.Auction, error) {
	if auction.ID == bson.NilObjectID {
		auction.ID = bson.NewObjectID()
	}

	return uc.AuctionRepo.Create(auction)
}
