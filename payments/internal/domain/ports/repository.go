package ports

import "payments/internal/domain/entity"

type PaymentRepository interface {
	Save(payment *entity.Payment) error
	FindByAuctionID(auctionID string) (*entity.Payment, error)
	UpdateStatus(auctionID string, status string) error
	FindExpired(duration int64) ([]*entity.Payment, error)
}
