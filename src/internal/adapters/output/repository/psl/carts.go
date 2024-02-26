package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
	log "github.com/sirupsen/logrus"
)

type Carts struct {
	Repository
}

func NewCarts(c *sql.DB, l *log.Logger) *Carts {
	return &Carts{
		Repository: Repository{
			client: c,
			log:    l,
		},
	}
}

func (p *Carts) GetAll(ctx context.Context, cartsID string) ([]*domain.ShoppingCart, error) {
	return nil, nil
}

func (p *Carts) Insert(ctx context.Context, carts *domain.ShoppingCart) error {
	p.log.Infoln("Insert", carts)
	query := "INSERT INTO ShoppingCarts(cartID, OwnerID) VALUES ($1, $2);"
	p.log.Debug(query)

	uid := uuid.New()
	values := []interface{}{
		uid.String(),
		&carts.OwnerID,
	}
	carts.ID = uid

	return p.insert(ctx, query, values)
}
