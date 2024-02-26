package usecase

import (
	"context"

	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
	"github.com/isaias-dgr/ecommerce_service/src/internal/core/ports"
	log "github.com/sirupsen/logrus"
)

type Products struct {
	l            *log.Logger
	productsRepo ports.ProductsRepo
}

func NewProducts(log *log.Logger, repo ports.ProductsRepo) *Products {
	return &Products{
		l:            log,
		productsRepo: repo,
	}
}

func (p Products) Create(ctx context.Context, product *domain.Product) error {
	p.l.Infof("Creating %v", product)
	err := p.productsRepo.Insert(ctx, product)
	if err != nil {
		return err
	}
	return nil
}
