package ports

import (
	"context"

	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
)

type PaymentsRepo interface {
	GetAllByUser(context.Context, string) ([]*domain.Payment, error)
	Insert(context.Context, *domain.Payment) error
}
