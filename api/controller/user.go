package controller

import (
	"net/http"

	"github.com/FastDefence/go_learn/api/model"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	name := c.QueryParam("name")

	user := model.User{
		Name: name,
	}
	if err := c.Bind(&user); err != nil {
		return err
	}
	model.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
	users := []model.User{}
	model.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	model.DB.Take(&user)
	return c.JSON(http.StatusOK, user)
}
