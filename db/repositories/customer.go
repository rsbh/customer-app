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

type CreateCustomersBody struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

func (repo *CustomerRepo) CreateCustomers(ctx context.Context, body CreateCustomersBody) (schema.Customer, error) {
	customer := schema.Customer{}
	query, _, err := goqu.Insert("customers").Rows(body).Returning("*").ToSQL()
	if err != nil {
		return schema.Customer{}, err
	}
	err = repo.DB.QueryRowxContext(ctx, query).StructScan(&customer)
	if err != nil {
		return schema.Customer{}, err
	}
	return customer, nil
}
