package repository

import (
	"context"

	"github.com/LehcimW4R/inventory-manager/internal/models"
	"github.com/jmoiron/sqlx"
)

// Repository is the interface that wraps the basic CRUD operations
//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	SaveUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
	GetUserRoles(ctx context.Context, userID int64) ([]models.UserRole, error)

	SaveProduct(ctx context.Context, name, description string, price float32, created_by int64) error
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProduct(ctx context.Context, id int64) (*models.Product, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
