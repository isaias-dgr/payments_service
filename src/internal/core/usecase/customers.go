package usecase

import (
	"context"

	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
	"github.com/isaias-dgr/ecommerce_service/src/internal/core/ports"
	log "github.com/sirupsen/logrus"
)

type Customers struct {
	l             *log.Logger
	customersRepo ports.CustomersRepo
}

func NewCustomers(log *log.Logger, repo ports.CustomersRepo) *Customers {
	return &Customers{
		l:             log,
		customersRepo: repo,
	}
}

func (p *Customers) Create(ctx context.Context, customer *domain.Customer) error {
	p.l.Infof("Creating %v", customer)
	err := p.customersRepo.Insert(ctx, customer)
	if err != nil {
		return err
	}
	return nil
}

func (p *Customers) GetAll(ctx context.Context) ([]*domain.Customer, error) {
	p.l.Infoln("Getting all customer")
	customers, err := p.customersRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return customers, nil
}
