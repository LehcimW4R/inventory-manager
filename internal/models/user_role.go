package models

type UserRole struct {
	ID     int64 `db:"id" json:"-"`
	UserID int64 `db:"user_id" json:"user_id"`
	RoleID int64 `db:"role_id" json:"role_id"`
}
