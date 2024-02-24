package domain

import (
	"fmt"

	"github.com/google/uuid"
)

// Individuals who make online purchases and complete payments
// through the platform.
type Customer struct {
	ID            uuid.UUID
	Name          string
	Email         string
	PaymentMethod PaymentMethod
}

func NewCustomer() *Customer {
	return &Customer{}
}

func (c Customer) Print() string {
	return fmt.Sprintf("Customer: %s", c.ID)
}
