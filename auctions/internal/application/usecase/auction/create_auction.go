package auction

import (
	"github.com/lioarce01/auction-microservices/internal/domain/entities"
	"github.com/lioarce01/auction-microservices/internal/domain/repositories"
	"github.com/lioarce01/auction-microservices/internal/infrastructure/services"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateAuctionUseCase struct {
	AuctionRepo repositories.AuctionRepository
	AuthService *services.AuthService
}

func NewCreateAuctionUseCase(repo repositories.AuctionRepository, authService *services.AuthService) *CreateAuctionUseCase {
	return &CreateAuctionUseCase{
		AuctionRepo: repo,
		AuthService: authService,
	}
}

func (uc *CreateAuctionUseCase) Execute(auction entities.Auction) (entities.Auction, error) {
	creatorIDStr := auction.CreatorID.Hex()

	creatorID, err := uc.AuthService.GetCreatorID(creatorIDStr)
	if err != nil {
		return entities.Auction{}, err
	}

	objectID, err := bson.ObjectIDFromHex(creatorID)
	if err != nil {
		return entities.Auction{}, err
	}
	auction.CreatorID = objectID

	return uc.AuctionRepo.Create(auction)
}
