package bid

import (
	"github.com/lioarce01/auction-microservices/internal/domain/entities"
	"github.com/lioarce01/auction-microservices/internal/domain/repositories"
	"github.com/lioarce01/auction-microservices/internal/infrastructure/services"
)

type CreateBidUseCase struct {
	BidRepo     repositories.BidRepository
	AuthService *services.AuthService
}

func NewCreateBidUseCase(repo repositories.BidRepository, authService *services.AuthService) *CreateBidUseCase {
	return &CreateBidUseCase{
		BidRepo:     repo,
		AuthService: authService,
	}
}

func (uc *CreateBidUseCase) Execute(token string, auctionID string, price float64) (entities.Bid, error) {
	creatorIDStr, err := uc.AuthService.GetCreatorID(token)
	if err != nil {
		return entities.Bid{}, err
	}

	createdBid, err := uc.BidRepo.CreateBid(auctionID, creatorIDStr, price)
	if err != nil {
		return entities.Bid{}, err
	}

	return createdBid, nil
}
