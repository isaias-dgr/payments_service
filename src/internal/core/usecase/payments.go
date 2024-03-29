package usecase

import (
	"context"

	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
	"github.com/isaias-dgr/ecommerce_service/src/internal/core/ports"
	log "github.com/sirupsen/logrus"
)

type Payments struct {
	l           *log.Logger
	paymentRepo ports.PaymentsRepo
}

func NewPayments(log *log.Logger, repo ports.PaymentsRepo) *Payments {
	return &Payments{
		l:           log,
		paymentRepo: repo,
	}
}

func (p Payments) Create(ctx context.Context, payment *domain.Payment) error {
	p.l.Infof("Creating %v", payment)
	err := p.paymentRepo.Insert(ctx, payment)
	if err != nil {
		return err
	}
	return nil
}

func (p Payments) Cancel() error {
	return nil
}

func (p Payments) Refund(ctx context.Context, userID, paymentID string) error {
	return nil
}

func (p Payments) GetAll(ctx context.Context, paymentID string) ([]*domain.Payment, error) {
	p.l.Infof("Getting all Payments for user %s", paymentID)
	payments, err := p.paymentRepo.GetAllByUser(ctx, paymentID)
	if err != nil {
		return nil, err
	}
	return payments, nil
}
