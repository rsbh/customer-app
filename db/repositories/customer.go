package repositories

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/rsbh/customer-app/db/schema"
)

type CustomerRepo struct {
	DB *sqlx.DB
}

func NewCustomerRepo(DB *sqlx.DB) *CustomerRepo {
	return &CustomerRepo{
		DB: DB,
	}
}

func (repo *CustomerRepo) ListCustomers(ctx context.Context) ([]schema.Customer, error) {
	customers := []schema.Customer{}
	query, _, err := goqu.Select("*").From("customers").ToSQL()
	if err != nil {
		return nil, err
	}
	if err = repo.DB.SelectContext(ctx, &customers, query); err != nil {
		return nil, err
	}
	return customers, nil
}
