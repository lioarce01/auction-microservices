package auction

import (
	"fmt"

	"github.com/lioarce01/auction-microservices/internal/domain/repositories"
	"github.com/lioarce01/auction-microservices/internal/infrastructure/services"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DeleteAuctionUseCase struct {
	AuctionRepo repositories.AuctionRepository
	AuthService *services.AuthService
}

func NewDeleteAuctionUseCase(repo repositories.AuctionRepository, authService *services.AuthService) *DeleteAuctionUseCase {
	return &DeleteAuctionUseCase{
		AuctionRepo: repo,
		AuthService: authService,
	}
}

func (uc *DeleteAuctionUseCase) Execute(token string, auctionID string) error {
	creatorID, err := uc.AuthService.GetCreatorID(token)
	if err != nil {
		fmt.Println("❌ [UseCase] Error obteniendo creatorID:", err)
		return err
	}

	err = uc.AuctionRepo.Delete(creatorID, auctionID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("unauthorized: auction not found or you are not the creator")
		}
		return err
	}

	return nil
}
