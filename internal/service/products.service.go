package service

import (
	context "context"
	"errors"

	models "github.com/LehcimW4R/inventory-manager/internal/models"
)

var validRolesToAddProduct []int64 = []int64{1, 2}
var ErrInvalidPermissions = errors.New("user does not have permission to add product")

// GetProduct implements Service.
func (s *serv) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	p, err := s.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}, nil
}

// GetProducts implements Service.
func (s *serv) GetProducts(ctx context.Context) ([]models.Product, error) {
	pp, err := s.repo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := []models.Product{}

	for _, item := range pp {
		products = append(products, models.Product{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
		})
	}

	return products, nil
}

// SaveProduct implements Service.
func (s *serv) AddProduct(ctx context.Context, product models.Product, email string) error {

	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	roles, err := s.repo.GetUserRoles(ctx, u.ID)
	if err != nil {
		return err
	}

	userCanAdd := false

	for _, item := range roles {
		for _, vr := range validRolesToAddProduct {
			if vr == item.RoleID {
				userCanAdd = true
			}
		}
	}

	if !userCanAdd {
		return ErrInvalidPermissions
	}

	if err := s.repo.SaveProduct(ctx, product.Name, product.Description, product.Price, u.ID); err != nil {
		return err
	}
	return nil
}
