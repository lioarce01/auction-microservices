package auction

import (
	"github.com/lioarce01/auction-microservices/internal/domain/entities"
	"github.com/lioarce01/auction-microservices/internal/domain/repositories"
)

type ListAuctionsUseCase struct {
	AuctionRepo repositories.AuctionRepository
}

func NewListAuctionsUseCase(repo repositories.AuctionRepository) *ListAuctionsUseCase {
	return &ListAuctionsUseCase{AuctionRepo: repo}
}

func (uc *ListAuctionsUseCase) Execute() ([]entities.Auction, error) {
	return uc.AuctionRepo.List()
}
