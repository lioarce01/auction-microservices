package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lioarce01/auction-microservices/internal/application/usecase/bid"
)

type BidHandler struct {
	CreateBidUseCase *bid.CreateBidUseCase
	ListBidsUseCase  *bid.ListBidsUseCase
}

func NewBidHandler(
	createBidUseCase *bid.CreateBidUseCase,
	listBidsUseCase *bid.ListBidsUseCase,
) *BidHandler {
	return &BidHandler{
		CreateBidUseCase: createBidUseCase,
		ListBidsUseCase:  listBidsUseCase,
	}
}

func (h *BidHandler) ListAllBids(c *gin.Context) {
	auctionID := c.Param("id")

	bids, err := h.ListBidsUseCase.Execute(auctionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bids)
}

func (h *BidHandler) CreateBid(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req struct {
		Price  float64 `json:"price" binding:"required"`
		UserID string  `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auctionID := c.Param("id")

	bid, err := h.CreateBidUseCase.Execute(token, auctionID, req.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, bid)
}
