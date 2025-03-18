package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lioarce01/auction-microservices/internal/application/usecase/auction"
	"github.com/lioarce01/auction-microservices/internal/domain/entities"
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
	var auction entities.Auction
	if err := c.ShouldBindJSON(&auction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auction, err := h.CreateAuctionUseCase.Execute(auction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, auction)
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
