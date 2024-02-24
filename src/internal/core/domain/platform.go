package domain

import (
	"fmt"

	"github.com/google/uuid"
)

// An application that validates requests, stores card information,
// and manages payment requests and responses to and from the acquiring bank.
type Platform struct {
	ID             uuid.UUID
	Name           string
	SupportedCards []string
	Merchants      []Merchant
}

func NewPlatform() *Platform {
	return &Platform{}
}

func (p Platform) Print() string {
	return fmt.Sprintf("Platform: %s", p.ID)
}
