package worker

import (
	"math/rand"
	"time"

	"github.com/fiorellizz/gopayment/internal/logger"
	"github.com/fiorellizz/gopayment/internal/payment"
)

// StartWorker inicia um worker que processa pagamentos da fila
func StartWorker(id int, jobs <-chan payment.Payment) {
	go func() {
		for job := range jobs {

			// Logando início do processamento
			logger.Log.Info("Worker started processing",
				"worker_id", id,
				"payment_id", job.ID,
			)

			// Simulando processamento
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

			// Logando fim do processamento
			logger.Log.Info("Payment processed successfully",
				"worker_id", id,
				"payment_id", job.ID,
				"amount", job.Amount,
			)
		}
	}()
}