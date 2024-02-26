package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
	log "github.com/sirupsen/logrus"
)

type Customers struct {
	Repository
}

func NewCustomers(c *sql.DB, l *log.Logger) *Customers {
	return &Customers{
		Repository: Repository{
			client: c,
			log:    l,
		},
	}
}

func (c *Customers) GetAll(ctx context.Context) ([]*domain.Customer, error) {
	c.log.Info("💾 Getting all customers")
	query := "SELECT CustomerID, Name, Email FROM Customers;"
	c.log.Debug(query)
	rows, err := c.client.QueryContext(ctx, query)
	if err != nil {
		c.log.Error(err.Error())
		return nil, errors.New("query_context")
	}
	defer rows.Close()

	customers := []*domain.Customer{}
	for rows.Next() {
		customer := &domain.Customer{}
		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Email)
		if err != nil {
			c.log.Error(err.Error())
			return nil, errors.New(err.Error())
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		c.log.Error(err.Error())
		return nil, errors.New("row_corrupt")
	}
	return customers, nil
}

func (c *Customers) Insert(ctx context.Context, customer *domain.Customer) error {
	c.log.Infoln("Insert", customer)
	query := "INSERT INTO Customers(CustomerID, Name, Email) VALUES ($1, $2, $3);"
	c.log.Debug(query)

	uid := uuid.New()
	values := []interface{}{
		uid.String(),
		&customer.Name,
		&customer.Email,
	}
	customer.ID = uid

	return c.insert(ctx, query, values)
}
