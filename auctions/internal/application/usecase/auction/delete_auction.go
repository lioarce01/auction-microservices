package auction

import "github.com/lioarce01/auction-microservices/internal/domain/repositories"

type DeleteAuctionUseCase struct {
	AuctionRepo repositories.AuctionRepository
}

func NewDeleteAuctionUseCase(repo repositories.AuctionRepository) *DeleteAuctionUseCase {
	return &DeleteAuctionUseCase{AuctionRepo: repo}
}

func (uc *DeleteAuctionUseCase) Execute(creatorID string, auctionID string) error {
	return uc.AuctionRepo.Delete(creatorID, auctionID)
}
