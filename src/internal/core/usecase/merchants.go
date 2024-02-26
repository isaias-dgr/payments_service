package usecase

import (
	"context"

	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
	"github.com/isaias-dgr/ecommerce_service/src/internal/core/ports"
	log "github.com/sirupsen/logrus"
)

type Merchants struct {
	l             *log.Logger
	merchantsRepo ports.MerchantsRepo
}

func NewMerchants(log *log.Logger, repo ports.MerchantsRepo) *Merchants {
	return &Merchants{
		l:             log,
		merchantsRepo: repo,
	}
}

func (c *Merchants) Create(ctx context.Context, merchant *domain.Merchant) error {
	c.l.Infof("Creating %s", merchant.BusinessID)
	err := c.merchantsRepo.Insert(ctx, merchant)
	if err != nil {
		return err
	}
	return nil
}
