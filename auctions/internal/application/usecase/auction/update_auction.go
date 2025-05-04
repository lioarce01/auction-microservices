package auction

import (
	"github.com/lioarce01/auction-microservices/internal/domain/entities"
	"github.com/lioarce01/auction-microservices/internal/domain/repositories"
)

type UpdateAuctionUseCase struct {
	AuctionRepo repositories.AuctionRepository
}

func NewUpdateAuctionUseCase(repo repositories.AuctionRepository) *UpdateAuctionUseCase {
	return &UpdateAuctionUseCase{AuctionRepo: repo}
}

func (uc *UpdateAuctionUseCase) Execute(creatorID string, auction entities.Auction) (entities.Auction, error) {
	return uc.AuctionRepo.Update(creatorID, auction)
}
