package worker

import (
	"context"
	"fmt"
	"time"

	"payments/internal/application/usecase"
	"payments/internal/domain/ports"
)

type ExpirationWorker struct {
	expirationUC *usecase.ExpirationUseCase
	ticker       *time.Ticker
}

func NewExpirationWorker(repo ports.PaymentRepository) *ExpirationWorker {
	// threshold de 48 horas
	expirationUC := usecase.NewExpirationUseCase(repo, 48*time.Hour)
	return &ExpirationWorker{
		expirationUC: expirationUC,
		ticker:       time.NewTicker(1 * time.Hour),
	}
}

func (w *ExpirationWorker) Start(ctx context.Context) {
	defer w.ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-w.ticker.C:
			if err := w.expirationUC.Execute(ctx); err != nil {
				// loguear el error
				fmt.Printf("error running expiration job: %v\n", err)
			}
		}
	}
}
