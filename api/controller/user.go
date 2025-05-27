package controller

import (
	"net/http"

	"github.com/FastDefence/go_learn/api/model"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	nema := c.queryParam("name")

	user := model.User{
		Name: name,
	}
	if err := c.Bind(&user); err != nil {
		return err
	}
	model.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}
