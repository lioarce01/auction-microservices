package usecase

import (
	"context"
	"fmt"
	"time"

	"payments/internal/domain/entity"
	"payments/internal/domain/ports"
)

type ExpirationUseCase struct {
	repo                ports.PaymentRepository
	expirationThreshold time.Duration
}

func NewExpirationUseCase(repo ports.PaymentRepository, threshold time.Duration) *ExpirationUseCase {
	return &ExpirationUseCase{
		repo:                repo,
		expirationThreshold: threshold,
	}
}

// Execute busca y expira pagos cuyo tiempo de creación excede el umbral
func (uc *ExpirationUseCase) Execute(ctx context.Context) error {
	// Convertimos el threshold a milisegundos para el repositorio
	ms := uc.expirationThreshold.Milliseconds()

	// 1. Obtener todos los pagos vencidos
	payments, err := uc.repo.FindExpired(ms)
	if err != nil {
		return fmt.Errorf("failed to find expired payments: %w", err)
	}

	// 2. Para cada pago, cambiar el estado a "expired"
	for _, p := range payments {
		p.Status = entity.StatusExpired
		if err := uc.repo.UpdateStatus(p.AuctionID, string(p.Status)); err != nil {
			// loguear el error y continuar con los demás
			fmt.Printf("error expiring payment %s: %v\n", p.ID, err)
		}
	}

	return nil
}
