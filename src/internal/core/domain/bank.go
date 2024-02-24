package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Facilitates the actual retrieval of funds from the customer's card
// and transfers them to the merchant.
type AcquiringBank struct {
	ID       uuid.UUID
	Name     string
	BankCode string
}

func NewBanck() *AcquiringBank {
	return &AcquiringBank{}
}

func (b AcquiringBank) Print() string {
	return fmt.Sprintf("Customer: %s", b.ID)
}

// PaymentMethod represents a customer's payment method details.
type PaymentMethod struct {
	CardNumber     string
	ExpirationDate time.Time
	CVV            string
	CardholderName string
}

// BankAccount represents a bank account information for a Merchant.
type BankAccount struct {
	AccountNumber string
	BankName      string
	BranchCode    string
}
