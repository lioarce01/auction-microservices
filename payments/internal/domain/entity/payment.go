package entity

import "time"

// Definimos un tipo para los estados de pago
type PaymentStatus string

const (
	StatusPending PaymentStatus = "pending"
	StatusPaid    PaymentStatus = "paid"
	StatusExpired PaymentStatus = "expired"
)

type Payment struct {
	ID          string
	AuctionID   string
	Amount      float64
	Status      PaymentStatus // ahora usa el tipo PaymentStatus
	PaymentLink string
	CreatedAt   time.Time
}

func NewPayment(auctionID string, amount float64) *Payment {
	return &Payment{
		AuctionID: auctionID,
		Amount:    amount,
		Status:    StatusPending,
		CreatedAt: time.Now().UTC(),
	}
}
