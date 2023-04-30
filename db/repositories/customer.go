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

const CUSTOMER_TABLE = "customers"

func (repo *CustomerRepo) ListCustomers(ctx context.Context) ([]schema.Customer, error) {
	customers := []schema.Customer{}
	query, _, err := goqu.Select("*").From(CUSTOMER_TABLE).ToSQL()
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
	query, _, err := goqu.Insert(CUSTOMER_TABLE).Rows(body).Returning("*").ToSQL()
	if err != nil {
		return schema.Customer{}, err
	}
	err = repo.DB.QueryRowxContext(ctx, query).StructScan(&customer)
	if err != nil {
		return schema.Customer{}, err
	}
	return customer, nil
}

func (repo *CustomerRepo) GetCustomer(ctx context.Context, id int) (schema.Customer, error) {
	customer := schema.Customer{}
	query, _, err := goqu.Select("*").From(CUSTOMER_TABLE).Where(goqu.Ex{
		"id": id,
	}).ToSQL()
	if err != nil {
		return schema.Customer{}, err
	}
	err = repo.DB.GetContext(ctx, &customer, query)
	if err != nil {
		return schema.Customer{}, err
	}
	return customer, nil
}

func (repo *CustomerRepo) UpdateCustomers(ctx context.Context, id int, body CreateCustomersBody) (schema.Customer, error) {
	customer := schema.Customer{}
	query, _, err := goqu.Update(CUSTOMER_TABLE).Set(body).Where(goqu.Ex{"id": id}).Returning("*").ToSQL()
	if err != nil {
		return schema.Customer{}, err
	}
	err = repo.DB.QueryRowxContext(ctx, query).StructScan(&customer)
	if err != nil {
		return schema.Customer{}, err
	}
	return customer, nil
}
