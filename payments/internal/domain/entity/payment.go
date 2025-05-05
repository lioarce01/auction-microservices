package entity

import "time"

type Payment struct {
	ID          string
	AuctionID   string
	Amount      float64
	Status      string // "pending", "paid", "expired"
	PaymentLink string
	CreatedAt   time.Time
}

func NewPayment(auctionID string, amount float64) *Payment {
	return &Payment{
		AuctionID: auctionID,
		Amount:    amount,
		Status:    "pending",
		CreatedAt: time.Now().UTC(),
	}
}
