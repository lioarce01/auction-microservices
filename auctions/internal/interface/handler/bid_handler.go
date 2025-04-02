package handler

import (
	"net/http"

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
	panic("not implemented")
}
