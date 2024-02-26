package domain

import (
	"fmt"

	"github.com/google/uuid"
)

// ShoppingCartItem represents a line item in a shopping cart.
type ShoppingCartItem struct {
	ProductID uuid.UUID
	Quantity  int
}

// ShoppingCart represents a user's shopping cart.
type ShoppingCart struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
	Items   []ShoppingCartItem
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}

func (s ShoppingCart) Print() string {
	return fmt.Sprintf("ShoppingCart: %s", s.ID)
}
