package gateway

import (
	"context"
	"fmt"
	"payments/internal/domain/ports"
	"strconv"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/order"
)

// MercadoPagoAdapter implementa ports.PaymentGateway
type MercadoPagoAdapter struct {
	client order.Client
}

// NewMercadoPagoAdapter crea una instancia del adaptador
func NewMercadoPagoAdapter(accessToken string) ports.PaymentGateway {
	cfg, err := config.New(accessToken)
	if err != nil {
		panic(fmt.Sprintf("failed to create MercadoPago config: %v", err))
	}
	return &MercadoPagoAdapter{
		client: order.NewClient(cfg),
	}
}

// CreatePayment crea un pago (boleto/ticket) y devuelve el link
func (a *MercadoPagoAdapter) CreatePayment(amount float64, description, auctionID string) (string, error) {
	ctx := context.Background()

	amountStr := strconv.FormatFloat(amount, 'f', 2, 64)

	req := order.Request{
		Type:              "online", // flujo de pago
		TotalAmount:       amountStr,
		ExternalReference: auctionID,
		Description:       description,
		// Opcional: si queres enviar email del payer:
		// Payer: &order.PayerRequest{Email: payerEmail},
		Transactions: &order.TransactionRequest{
			Payments: []order.PaymentRequest{
				{
					Amount: amountStr,
					PaymentMethod: &order.PaymentMethodRequest{
						Type: "ticket", // pago por boleto/voucher
					},
				},
			},
		},
	}

	resource, err := a.client.Create(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to create MercadoPago order: %w", err)
	}

	if tx := resource.Transactions; len(tx.Payments) > 0 {
		pm := tx.Payments[0].PaymentMethod
		// Solo comprobamos la cadena; el struct nunca es nil
		if pm.TicketURL != "" {
			return pm.TicketURL, nil
		}
	}

	return resource.ID, nil

}

func (a *MercadoPagoAdapter) GetPaymentStatus(orderID string) (string, error) {
	ctx := context.Background()
	resource, err := a.client.Get(ctx, orderID)
	if err != nil {
		return "", fmt.Errorf("failed to get order status: %w", err)
	}

	// Devolvemos el estado del primer pago si existe
	if tx := resource.Transactions; len(tx.Payments) > 0 {
		return string(tx.Payments[0].Status), nil
	}
	// Si no hay pagos, devolvemos el estado general de la orden
	return string(resource.Status), nil
}
