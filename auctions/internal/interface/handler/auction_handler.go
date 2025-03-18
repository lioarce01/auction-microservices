package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lioarce01/auction-microservices/internal/application/usecase/auction"
	"github.com/lioarce01/auction-microservices/internal/domain/entities"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type AuctionHandler struct {
	CreateAuctionUseCase *auction.CreateAuctionUseCase
	ListAuctionsUseCase  *auction.ListAuctionsUseCase
	GetOneAuctionUseCase *auction.GetOneAuctionUseCase
	UpdateAuctionUseCase *auction.UpdateAuctionUseCase
	DeleteAuctionUseCase *auction.DeleteAuctionUseCase
}

func NewAuctionHandler(
	listAuctionsUseCase *auction.ListAuctionsUseCase,
	getOneAuctionUseCase *auction.GetOneAuctionUseCase,
	createAuctionUseCase *auction.CreateAuctionUseCase,
	updateAuctionUseCase *auction.UpdateAuctionUseCase,
	deleteAuctionUseCase *auction.DeleteAuctionUseCase,
) *AuctionHandler {
	return &AuctionHandler{
		ListAuctionsUseCase:  listAuctionsUseCase,
		GetOneAuctionUseCase: getOneAuctionUseCase,
		CreateAuctionUseCase: createAuctionUseCase,
		UpdateAuctionUseCase: updateAuctionUseCase,
		DeleteAuctionUseCase: deleteAuctionUseCase,
	}
}

func (h *AuctionHandler) ListAllAuctions(c *gin.Context) {
	auctions, err := h.ListAuctionsUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auctions)
}

func (h *AuctionHandler) GetAuction(c *gin.Context) {
	id := c.Param("id")

	auction, err := h.GetOneAuctionUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auction)
}

func (h *AuctionHandler) CreateAuction(c *gin.Context) {
	var req entities.AuctionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userClaimsMap, ok := userClaims.(map[string]interface{})
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user claims"})
		return
	}

	creatorIDStr, ok := userClaimsMap["sub"].(string)
	if !ok || creatorIDStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ID de usuario no disponible en el token"})
		return
	}

	creatorID, err := bson.ObjectIDFromHex(creatorIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid creator ID format"})
		return
	}

	auction := entities.Auction{
		Title:        req.Title,
		Description:  req.Description,
		CreatorID:    creatorID,
		CurrentPrice: req.CurrentPrice,
		EndDate:      req.EndDate,
		Status:       req.Status,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	createdAuction, err := h.CreateAuctionUseCase.Execute(auction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdAuction)
}

func (h *AuctionHandler) UpdateAuction(c *gin.Context) {
	var auction entities.Auction
	id := c.GetString("id")

	if err := c.ShouldBindJSON(&auction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auction, err := h.UpdateAuctionUseCase.Execute(id, auction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auction)
}

func (h *AuctionHandler) DeleteAuction(c *gin.Context) {
	id := c.Param("id")

	creatorID := c.GetString("userID")

	err := h.DeleteAuctionUseCase.Execute(creatorID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Auction deleted successfully"})
}
