package repository

import (
	context "context"

	models "github.com/LehcimW4R/inventory-manager/internal/models"
)

const (
	qryInsertProduct = `
		insert into PRODUCTS (name, description, price, created_by) values (?,?,?,?);
	`
	qryFindProduct = `
		select 
			name, 
			description, 
			price, 
			created_by 
		from PRODUCTS 
		where id = ?;
	`
	qryFindAllProducts = `
		select name, description, price, created_by from PRODUCTS;
	`
)

// SaveProduct implements Repository.
func (r *repo) SaveProduct(ctx context.Context, name string, description string, price float32, created_by int64) error {
	_, err := r.db.ExecContext(ctx, name, description, price, created_by)
	return err
}

// GetProduct implements Repository.
func (r *repo) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	p := &models.Product{}
	err := r.db.GetContext(ctx, p, qryFindProduct, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// GetProducts implements Repository.
func (r *repo) GetProducts(ctx context.Context) ([]models.Product, error) {
	pp := []models.Product{}
	err := r.db.SelectContext(ctx, &pp, qryFindAllProducts)
	if err != nil {
		return nil, err
	}
	return pp, nil
}
