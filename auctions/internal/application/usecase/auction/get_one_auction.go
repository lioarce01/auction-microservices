package auction

import (
	"github.com/lioarce01/auction-microservices/internal/domain/entities"
	"github.com/lioarce01/auction-microservices/internal/domain/repositories"
)

type GetOneAuctionUseCase struct {
	AuctionRepo repositories.AuctionRepository
}

func NewGetOneAuctionUseCase(repo repositories.AuctionRepository) *GetOneAuctionUseCase {
	return &GetOneAuctionUseCase{AuctionRepo: repo}
}

func (uc *GetOneAuctionUseCase) Execute(id string) (entities.Auction, error) {
	return uc.AuctionRepo.GetOne(id)
}
