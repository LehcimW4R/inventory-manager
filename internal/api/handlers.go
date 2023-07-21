package api

import (
	"log"
	"net/http"

	"github.com/LehcimW4R/inventory-manager/internal/api/dtos"
	"github.com/LehcimW4R/inventory-manager/internal/service"
	"github.com/LehcimW4R/inventory-manager/security"
	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, params.Email, params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "user already exists"})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusCreated, nil)
}

func (a *API) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.LoginUser{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	u, err := a.serv.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	//TODO: create JWT token
	token, err := security.SignedLoginToken(u)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	//TODO: return JWT token
	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
	}

	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]string{"success": "true"})
}
