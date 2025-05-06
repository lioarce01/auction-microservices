package usecase

import (
	"payments/internal/domain/entity"
	"payments/internal/domain/ports"
)

type PaymentUseCase struct {
	repo    ports.PaymentRepository
	gateway ports.PaymentGateway
}

func NewPaymentUseCase(repo ports.PaymentRepository, gateway ports.PaymentGateway) *PaymentUseCase {
	return &PaymentUseCase{repo: repo, gateway: gateway}
}

func (uc *PaymentUseCase) CreatePayment(auctionID string, amount float64) (*entity.Payment, error) {
	payment := entity.NewPayment(auctionID, amount)

	paymentLink, err := uc.gateway.CreatePayment(amount, "Auction Payment", auctionID)
	if err != nil {
		return nil, err
	}
	payment.PaymentLink = paymentLink

	if err := uc.repo.Save(payment); err != nil {
		return nil, err
	}

	return payment, nil
}

func (uc *PaymentUseCase) HandleWebhook(paymentID string) error {
	status, err := uc.gateway.GetPaymentStatus(paymentID)
	if err != nil {
		return err
	}

	payment, err := uc.repo.FindByAuctionID(paymentID)
	if err != nil {
		return err
	}

	payment.Status = entity.PaymentStatus(status)
	return uc.repo.Save(payment)
}
