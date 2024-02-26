package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
	log "github.com/sirupsen/logrus"
)

type Merchants struct {
	Repository
}

func NewMerchants(c *sql.DB, l *log.Logger) *Merchants {
	return &Merchants{
		Repository: Repository{
			client: c,
			log:    l,
		},
	}
}

func (p *Merchants) GetAll(ctx context.Context, MerchantsID string) ([]*domain.ShoppingCart, error) {
	return nil, nil
}

func (p *Merchants) Insert(ctx context.Context, merchant *domain.Merchant) error {
	p.log.Infoln("Insert", merchant)
	query := "INSERT INTO ShoppingMerchants(MerchantID, Name, BusinessID) VALUES ($1, $2, $3);"
	p.log.Debug(query)

	uid := uuid.New()
	values := []interface{}{
		uid.String(),
		&merchant.Name,
		&merchant.BusinessID,
	}
	merchant.ID = uid
	return p.insert(ctx, query, values)
}
