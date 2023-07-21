package service

import (
	"context"
	"errors"

	"github.com/LehcimW4R/inventory-manager/internal/models"
	"github.com/LehcimW4R/inventory-manager/utils"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrRoleAlreadyAdded   = errors.New("role was already added for this user")
	ErrRoleNotFound       = errors.New("role not found for this user")
)

// LoginUser implements Service.
func (s *serv) LoginUser(ctx context.Context, email string, password string) (*models.User, error) {

	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	//TODO decript password
	bb, err := utils.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}
	decryptedPassword, err := utils.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	return u, nil
}

// RegisterUser implements Service.
func (s *serv) RegisterUser(ctx context.Context, email string, name string, password string) error {

	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	bb, err := utils.Encrypt([]byte(password))
	if err != nil {
		return err
	}
	encryptedPassword := utils.ToBase64(bb)

	return s.repo.SaveUser(ctx, email, name, encryptedPassword)
}

// AddUserRole implements Service.
func (s *serv) AddUserRole(ctx context.Context, userID int64, roleID int64) error {
	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}

	for _, item := range roles {
		if item.RoleID == roleID {
			return ErrRoleAlreadyAdded
		}
	}

	return s.repo.SaveUserRole(ctx, userID, roleID)
}

// RemoveUserRole implements Service.
func (s *serv) RemoveUserRole(ctx context.Context, userID int64, roleID int64) error {

	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}

	roleFound := false
	for _, item := range roles {
		if item.RoleID == roleID {
			roleFound = true
			break
		}
	}

	if !roleFound {
		return ErrRoleNotFound
	}

	return s.repo.RemoveUserRole(ctx, userID, roleID)
}
