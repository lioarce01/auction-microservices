package bid

import (
	"github.com/lioarce01/auction-microservices/internal/domain/entities"
	"github.com/lioarce01/auction-microservices/internal/domain/repositories"
)

type ListBidsUseCase struct {
	BidRepo repositories.BidRepository
}

func NewListBidsUseCase(repo repositories.BidRepository) *ListBidsUseCase {
	return &ListBidsUseCase{BidRepo: repo}
}

func (uc *ListBidsUseCase) Execute(auctionID string) ([]entities.Bid, error) {
	return uc.BidRepo.GetBids(auctionID)
}
