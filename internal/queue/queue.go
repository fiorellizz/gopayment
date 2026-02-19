package queue

import "github.com/fiorellizz/gopayment/internal/payment"

// PaymentQueue é a estrutura que representa a fila de pagamentos
type PaymentQueue struct {
	Jobs chan payment.Payment
}

// NewPaymentQueue cria uma nova fila de pagamentos com um buffer especificado
func NewPaymentQueue(buffer int) *PaymentQueue {
	return &PaymentQueue{
		Jobs: make(chan payment.Payment, buffer),
	}
}