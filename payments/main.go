package main

import (
	"context"
	"log"
	"payments/config"
	"payments/internal/application/usecase"
	"payments/internal/infrastructure/gateway"
	"payments/internal/infrastructure/repository"
	"payments/internal/interfaces/handlers"
	"payments/internal/interfaces/worker"
	servicediscovery "payments/internal/service-discovery"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.Load()

	// Initialize Consul service discovery
	consulService, err := servicediscovery.NewConsulService()
	if err != nil {
		log.Fatal("error initializing consul:", err)
	}

	err = consulService.RegisterService("auctions", "auctions", "auctions", 8080)
	if err != nil {
		log.Fatalf("error registering consul service: %v", err)
	}

	// Infrastructure layer
	db := repository.NewPostgresDB(cfg.DatabaseURL)
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go expirationWorker.Start(ctx)

	r.Run(":4001")
}
