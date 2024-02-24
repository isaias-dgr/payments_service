package domain

import (
	"fmt"

	"github.com/google/uuid"
)

// Product represents an item that can be purchased.
type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       float64
	Stock       int
}

func NewProduct() *Product {
	return &Product{}
}

func (p Product) Print() string {
	return fmt.Sprintf("Product: %s", p.ID)
}
