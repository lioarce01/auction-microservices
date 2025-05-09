package ports

type PaymentGateway interface {
	CreatePayment(amount float64, description string, auctionID string) (string, error)
	GetPaymentStatus(paymentID string) (string, error)
}
