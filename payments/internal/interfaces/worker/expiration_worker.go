package worker

import (
	"payments/internal/domain/ports"
	"time"
)

type ExpirationWorker struct {
	paymentRepo ports.PaymentRepository
}

func NewExpirationWorker(repo ports.PaymentRepository) *ExpirationWorker {
	return &ExpirationWorker{paymentRepo: repo}
}

func (w *ExpirationWorker) Start() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		// Get payments older than 48 hours
		payments, _ := w.paymentRepo.FindExpired(int64((48 * time.Hour).Milliseconds()))

		for _, payment := range payments {
			_ = w.paymentRepo.UpdateStatus(payment.AuctionID, "expired")
		}
	}
}
