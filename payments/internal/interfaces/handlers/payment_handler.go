package handlers

import (
	"net/http"
	"payments/internal/application/usecase"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentUC *usecase.PaymentUseCase
}

func NewPaymentHandler(paymentUC *usecase.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{paymentUC: paymentUC}
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req struct {
		AuctionID string  `json:"auction_id"`
		Amount    float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err := h.paymentUC.CreatePayment(req.AuctionID, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Payment creation failed"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) HandleWebhook(c *gin.Context) {
	paymentID := c.Query("id")
	if err := h.paymentUC.HandleWebhook(paymentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Webhook processing failed"})
		return
	}
	c.Status(http.StatusOK)
}
