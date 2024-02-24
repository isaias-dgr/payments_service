package domain

import (
	"fmt"

	"github.com/google/uuid"
)

// The seller who utilizes the payment platform to receive payments from customers.
type Merchant struct {
	ID          uuid.UUID
	Name        string
	BusinessID  string
	AccountInfo BankAccount
}

func NewMerchant() *Merchant {
	return &Merchant{}
}

func (m Merchant) Print() string {
	return fmt.Sprintf("Merchant: %s", m.ID)
}
