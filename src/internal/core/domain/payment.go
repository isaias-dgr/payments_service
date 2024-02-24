package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Payment struct {
	ID              uuid.UUID
	PaymentMethodID uuid.UUID
	CartID          uuid.UUID
	Status          string
	Total           float64
}

func NewPayment() *Payment {
	return &Payment{}
}

func (p Payment) Print() string {
	return fmt.Sprintf("Payment: %s, Cart: %s", p.ID, p.CartID)
}
