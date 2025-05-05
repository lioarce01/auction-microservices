package main

import (
	"payments/config"
	"payments/internal/application/usecase"
	"payments/internal/infrastructure/gateway"
	"payments/internal/infrastructure/repository"
	"payments/internal/interfaces/handlers"
	"payments/internal/interfaces/worker"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.Load()

	// Infrastructure layer
	db := repository.NewPostgresDB(cfg.DB)
	paymentRepo := repository.NewPostgresPaymentRepo(db)
	mpGateway := gateway.NewMercadoPagoAdapter(cfg.MercadoPagoToken)

	// Application layer
	paymentUC := usecase.NewPaymentUseCase(paymentRepo, mpGateway)

	// Interface layer
	paymentHandler := handlers.NewPaymentHandler(paymentUC)

	// Gin setup
	r := gin.Default()
	r.POST("/payments", paymentHandler.CreatePayment)
	r.POST("/webhook", paymentHandler.HandleWebhook)

	// Start expiration worker
	expirationWorker := worker.NewExpirationWorker(paymentRepo)
	go expirationWorker.Start()

	r.Run(":4001")
}
