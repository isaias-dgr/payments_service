package repositories

import (
	"context"
	"database/sql"
	"errors"

	// "github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
	log "github.com/sirupsen/logrus"
)

type Repository struct {
	client *sql.DB
	log    *log.Logger
}

func (r *Repository) insert(ctx context.Context, query string, values []interface{}) error {
	stmt, err := r.client.PrepareContext(ctx, query)
	if err != nil {
		r.log.Error(err.Error())
		return errors.New("query_prepare_ctx")
	}

	res, err := stmt.ExecContext(ctx, values...)
	if err != nil {
		r.log.Error(err.Error())
		return errors.New("query_exec")
	}

	affect, err := res.RowsAffected()
	if err != nil {
		r.log.Error(err.Error())
		return errors.New("query_exec")
	}
	r.log.Infof("%d saved", affect)
	return nil
}
