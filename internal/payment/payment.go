package payment

// Payment representa um pagamento a ser processado
type Payment struct {
	ID     string
	Amount float64
	UserID string
}