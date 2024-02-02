package repositories

import (
	"context"
	"database/sql"
	"ms-order/domain/entity"
)

type Order interface {
	Create(ctx context.Context, order entity.Order) (int64, error)
	FindAll(ctx context.Context) ([]entity.Order, error)
	FindById(ctx context.Context, id int) (entity.Order, error)
}

type OrderImpl struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) Order {
	return &OrderImpl{
		DB: db,
	}
}

func (p *OrderImpl) Create(ctx context.Context, order entity.Order) (int64, error) {
	query := `INSERT INTO orders (customer_id, product_id, quantity, total) VALUES ($1, $2, $3, $4)`

	result, err := p.DB.ExecContext(ctx, query, order.CustomerID, order.ProductID, order.Qty, order.Total)
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func (p *OrderImpl) FindAll(ctx context.Context) ([]entity.Order, error) {
	query := `SELECT id, customer_id, product_id, quantity, total, created_at, updated_at FROM orders`

	row, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	order := []entity.Order{}
	for row.Next() {
		o := entity.Order{}

		if err := row.Scan(&o.ID, o.CustomerID, o.ProductID, o.Qty, o.Total, o.CreatedAt, o.UpdatedAt); err != nil {
			return order, err
		}

		order = append(order, o)
	}

	return order, nil
}

func (p *OrderImpl) FindById(ctx context.Context, id int) (entity.Order, error) {
	order := entity.Order{}

	query := `SELECT id, customer_id, product_id, quantity, total, created_at, updated_at FROM order WHERE id = $1`

	row, err := p.DB.QueryContext(ctx, query, id)
	if err != nil {
		return order, err
	}

	for row.Next() {
		if err := row.Scan(&order.ID, &order.CustomerID, &order.ProductID, &order.Qty, &order.Total, &order.CreatedAt, &order.UpdatedAt); err != nil {
			return order, err
		}
	}

	return order, nil
}
