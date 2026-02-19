package main

import (
	"fmt"
	"time"

	"github.com/fiorellizz/gopayment/internal/logger"
	"github.com/fiorellizz/gopayment/internal/payment"
	"github.com/fiorellizz/gopayment/internal/queue"
	"github.com/fiorellizz/gopayment/internal/worker"
)

func main() {

	// Inicializando logger
	logger.Init()

	// Criando fila com capacidade de 10 pagamentos
	paymentQueue := queue.NewPaymentQueue(10)

	// Criando 3 workers
	for i := 1; i <= 3; i++ {
		worker.StartWorker(i, paymentQueue.Jobs)
	}

	// Simulando recebimento de pagamentos
	for i := 1; i <= 10; i++ {
		p := payment.Payment{
			ID:     fmt.Sprintf("pay_%d", i),
			Amount: float64(i * 100),
			UserID: fmt.Sprintf("user_%d", i),
		}

		logger.Log.Info("New payment received",
			"payment_id", p.ID,
		)

		paymentQueue.Jobs <- p
	}

	// Espera para workers processarem
	time.Sleep(10 * time.Second)
}