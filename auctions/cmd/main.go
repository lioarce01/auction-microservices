package main

import (
	"log"

	"github.com/lioarce01/auction-microservices/internal/interface/http"
)

func main() {
	r := http.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
