package repositories

import (
	"context"
	"database/sql"
	"ms-customer/domain/entity"
)

type Customer interface {
	Create(ctx context.Context, customer entity.Customer) error
	Update(ctx context.Context, customer entity.Customer) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]entity.Customer, error)
	FindById(ctx context.Context, id int) (entity.Customer, error)
}

type CustomerImpl struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) Customer {
	return &CustomerImpl{
		DB: db,
	}
}

func (p *CustomerImpl) Create(ctx context.Context, customer entity.Customer) error {
	query := `INSERT INTO customer (id, name, email) VALUES ($1, $2, $3)`

	_, err := p.DB.ExecContext(ctx, query, customer.ID, customer.Name, customer.Email)
	if err != nil {
		return err
	}

	return nil
}

func (p *CustomerImpl) Update(ctx context.Context, customer entity.Customer) error {
	query := `UPDATE customer SET name = $1, email = $2 WHERE id = $3`

	_, err := p.DB.ExecContext(ctx, query, customer.Name, customer.Email, customer.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *CustomerImpl) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM customer WHERE id = $1`

	_, err := p.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *CustomerImpl) FindAll(ctx context.Context) ([]entity.Customer, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM customer`

	row, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	customer := []entity.Customer{}
	for row.Next() {
		c := entity.Customer{}

		if err := row.Scan(&c.ID, &c.Name, &c.Email, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return customer, err
		}

		customer = append(customer, c)
	}

	return customer, nil
}

func (p *CustomerImpl) FindById(ctx context.Context, id int) (entity.Customer, error) {
	customer := entity.Customer{}

	query := `SELECT id, name, email, created_at, updated_at FROM customer WHERE id = $1`

	row, err := p.DB.QueryContext(ctx, query, id)
	if err != nil {
		return customer, err
	}

	for row.Next() {
		if err := row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt); err != nil {
			return customer, err
		}
	}

	return customer, nil
}
