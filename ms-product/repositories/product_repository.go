package repositories

import (
	"context"
	"database/sql"
	"ms-product/domain/entity"
)

type Product interface {
	Create(ctx context.Context, product entity.Product) error
	Update(ctx context.Context, product entity.Product) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]entity.Product, error)
	FindById(ctx context.Context, id int) (entity.Product, error)
}

type ProductImpl struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) Product {
	return &ProductImpl{
		DB: db,
	}
}

func (p *ProductImpl) Create(ctx context.Context, product entity.Product) error {
	query := `INSERT INTO product (id, name, price, stock) VALUES ($1, $2, $3, $4)`

	_, err := p.DB.ExecContext(ctx, query, product.ID, product.Name, product.Price, product.Stock)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductImpl) Update(ctx context.Context, product entity.Product) error {
	query := `UPDATE product SET name = $1, price = $2, stock = $3 WHERE id = $4`

	_, err := p.DB.ExecContext(ctx, query, product.Name, product.Price, product.Stock, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductImpl) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM product WHERE id = $1`

	_, err := p.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductImpl) FindAll(ctx context.Context) ([]entity.Product, error) {
	query := `SELECT id, name, price, stock, created_at, updated_at FROM product`

	row, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	product := []entity.Product{}
	for row.Next() {
		p := entity.Product{}

		if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return product, err
		}

		product = append(product, p)
	}

	return product, nil
}

func (p *ProductImpl) FindById(ctx context.Context, id int) (entity.Product, error) {
	product := entity.Product{}

	query := `SELECT id, name, price, stock, created_at, updated_at FROM product WHERE id = $1`

	row, err := p.DB.QueryContext(ctx, query, id)
	if err != nil {
		return product, err
	}

	for row.Next() {
		if err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return product, err
		}
	}

	return product, nil
}
