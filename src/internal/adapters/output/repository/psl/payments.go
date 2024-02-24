package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
	log "github.com/sirupsen/logrus"
)

type Payments struct {
	client *sql.DB
	log    *log.Logger
}

func NewPayments(c *sql.DB, l *log.Logger) *Payments {
	return &Payments{
		client: c,
		log:    l,
	}
}

func (p *Payments) fetch(
	ctx context.Context,
	stmt string,
	filters []interface{}) (ts []*domain.Payment, err error) {

	query := `SELECT PaymentID, PaymentMethodID, CartID, Total, Status FROM Payments ` + stmt
	rows, err := p.client.QueryContext(ctx, query)
	if err != nil {
		p.log.Error(err.Error())
		return nil, errors.New("query_context")
	}
	defer rows.Close()

	payments := []*domain.Payment{}
	for rows.Next() {
		payment := &domain.Payment{}
		err := rows.Scan(
			&payment.ID,
			&payment.PaymentMethodID,
			&payment.CartID,
			&payment.Total,
			&payment.Status)
		if err != nil {
			p.log.Error(err.Error())
			return nil, errors.New(err.Error())
		}
		payments = append(payments, payment)
	}
	if err := rows.Err(); err != nil {
		p.log.Error(err.Error())
		return nil, errors.New("row_corrupt")
	}

	return payments, nil
}

func (p *Payments) GetAllByUser(ctx context.Context, userID string) (t []*domain.Payment, err error) {
	p.log.Info("💾 Getting all payments")
	query := "WHERE PaymentID = $1"
	args := []interface{}{userID}

	payments, err := p.fetch(ctx, query, args)
	if err != nil {
		p.log.Info(err.Error())
		return nil, err
	}

	if len(payments) == 0 {
		p.log.Infof("Records %s Not Found", userID)
		return nil, errors.New("not_found")
	}

	return payments, nil
}

func (p *Payments) Insert(ctx context.Context, payment domain.Payment) (err error) {
	p.log.Infoln("Insert", payment)
	query := "INSERT INTO Payments(PaymentMethodID, CartID, Total, Status) VALUES ($1, $2, $3, $4);"
	fmt.Println(query)

	stmt, err := p.client.PrepareContext(ctx, query)
	if err != nil {
		p.log.Error(err.Error())
		return errors.New("query_prepare_ctx")
	}

	values := []interface{}{
		&payment.PaymentMethodID,
		&payment.CartID,
		&payment.Total,
		&payment.Status,
	}
	res, err := stmt.ExecContext(ctx, values...)
	if err != nil {
		p.log.Error(err.Error())
		return errors.New("query_exec")
	}
	affect, err := res.RowsAffected()
	if err != nil {
		p.log.Error(err.Error())
		return errors.New("query_exec")
	}
	p.log.Infof("%d Payment saved", affect)
	return nil
}
