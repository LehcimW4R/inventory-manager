package security

import (
	"github.com/LehcimW4R/inventory-manager/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

const key = "01234567890123456789012345678901"

func SignedLoginToken(u *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"name":  u.Name,
	})

	return token.SignedString([]byte(key))
}
