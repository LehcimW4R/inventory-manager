package repository

import (
	"context"

	"github.com/LehcimW4R/inventory-manager/internal/models"
)

const (
	qryInsertUser = `
		insert into USERS (email, name, password)
		values (?, ?, ?);
	`
	qryFindUserEmail = `
		select 
			id, 
			email, 
			name, 
			password 
		from USERS 
		where email = ?;
	`
	qryInsertUserRole = `
		insert into USER_ROLES (user_id, role_id)
		values (:user_id, :role_id);
	`
	qryRemoveUserRole = `
		delete from USER_ROLES 
		where user_id = :user_id and role_id = :role_id;
	`
	qryFindUserRolesByUserID = `
		select 
			user_id, 
			role_id
		from USER_ROLES 
		where user_id = ?;
	`
)

// SaveUser implements Repository.
func (r *repo) SaveUser(ctx context.Context, email string, name string, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

// GetUserByEmail implements Repository.
func (r *repo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	u := &models.User{}
	err := r.db.GetContext(ctx, u, qryFindUserEmail, email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// RemoveUserRole implements Repository.
func (r *repo) RemoveUserRole(ctx context.Context, userID int64, roleID int64) error {
	data := models.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	_, err := r.db.NamedExecContext(ctx, qryRemoveUserRole, data)

	return err
}

// SaveUserRole implements Repository.
func (r *repo) SaveUserRole(ctx context.Context, userID int64, roleID int64) error {
	data := models.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	_, err := r.db.NamedExecContext(ctx, qryInsertUserRole, data)

	return err
}

// GetUserRoles implements Repository.
func (r *repo) GetUserRoles(ctx context.Context, userID int64) ([]models.UserRole, error) {
	roles := []models.UserRole{}

	err := r.db.SelectContext(ctx, &roles, qryFindUserRolesByUserID, userID)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
