package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
	log "github.com/sirupsen/logrus"
)

type Products struct {
	Repository
}

func NewProducts(c *sql.DB, l *log.Logger) *Products {
	return &Products{
		Repository: Repository{
			client: c,
			log:    l,
		},
	}
}

func (p *Products) GetAll(ctx context.Context, productID string) ([]*domain.Product, error) {
	return nil, nil
}

func (p *Products) Insert(ctx context.Context, product *domain.Product) error {
	p.log.Infoln("Insert", product)
	query := "INSERT INTO Products(ProductID, Name, Price, Description, StockQuantity) VALUES ($1, $2, $3, $4, $5);"
	p.log.Debug(query)

	uid := uuid.New()
	values := []interface{}{
		uid.String(),
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
	}
	product.ID = uid

	return p.insert(ctx, query, values)
}
