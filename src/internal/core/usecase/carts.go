package usecase

import (
	"context"

	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
	"github.com/isaias-dgr/ecommerce_service/src/internal/core/ports"
	log "github.com/sirupsen/logrus"
)

type Carts struct {
	l         *log.Logger
	cartsRepo ports.CartsRepo
}

func NewCarts(log *log.Logger, repo ports.CartsRepo) *Carts {
	return &Carts{
		l:         log,
		cartsRepo: repo,
	}
}

func (c *Carts) Create(ctx context.Context, cart *domain.ShoppingCart) error {
	c.l.Infof("Creating %v", cart)
	err := c.cartsRepo.Insert(ctx, cart)
	if err != nil {
		return err
	}
	return nil
}
