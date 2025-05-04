package main

import (
	"log"

	"github.com/lioarce01/auction-microservices/internal/interface/http"
	servicediscovery "github.com/lioarce01/auction-microservices/internal/service-discovery"
)

func main() {
	r := http.SetupRouter()

	consulService, err := servicediscovery.NewConsulService()
	if err != nil {
		log.Fatal("error initializing consul:", err)
	}

	err = consulService.RegisterService("auctions", "auctions", "auctions", 8080)
	if err != nil {
		log.Fatalf("error registering consul service: %v", err)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
