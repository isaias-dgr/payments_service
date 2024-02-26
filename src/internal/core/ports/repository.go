package ports

import (
	"context"

	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
)

type PaymentsRepo interface {
	GetAllByUser(context.Context, string) ([]*domain.Payment, error)
	Insert(context.Context, *domain.Payment) error
}

type CartsRepo interface {
	Insert(context.Context, *domain.ShoppingCart) error
}

type CustomersRepo interface {
	GetAll(context.Context) ([]*domain.Customer, error)
	Insert(context.Context, *domain.Customer) error
}

type MerchantsRepo interface {
	Insert(context.Context, *domain.Merchant) error
}

type ProductsRepo interface {
	Insert(context.Context, *domain.Product) error
}
