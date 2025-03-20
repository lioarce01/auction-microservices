package auction

import (
	"fmt"

	"github.com/lioarce01/auction-microservices/internal/domain/entities"
	"github.com/lioarce01/auction-microservices/internal/domain/repositories"
	"github.com/lioarce01/auction-microservices/internal/infrastructure/services"
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

func (uc *CreateAuctionUseCase) Execute(token string, auction entities.Auction) (entities.Auction, error) {
	fmt.Println("🔍 [UseCase] Ejecutando creación de subasta")

	creatorIDStr, err := uc.AuthService.GetCreatorID(token)
	if err != nil {
		fmt.Println("❌ [UseCase] Error obteniendo creatorID:", err)
		return entities.Auction{}, err
	}
	fmt.Println("✅ [UseCase] CreatorID obtenido:", creatorIDStr)

	auction.CreatorID = creatorIDStr
	fmt.Println("📌 [UseCase] Subasta con CreatorID asignado:", auction)

	createdAuction, err := uc.AuctionRepo.Create(auction)
	if err != nil {
		fmt.Println("❌ [UseCase] Error guardando subasta en MongoDB:", err)
		return entities.Auction{}, err
	}

	fmt.Println("✅ [UseCase] Subasta creada:", createdAuction)
	return createdAuction, nil
}
