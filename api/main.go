package main

import (
	"net/http"

	_ "github.com/FastDefence/go_learn/api/docs" // docsを読み込む
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"

	"github.com/FastDefence/go_learn/api/controller"
	"github.com/FastDefence/go_learn/api/model"
)

func connect(c echo.Context) error {
	db, _ := model.DB.DB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "DB接続失敗しました")
	} else {
		return c.String(http.StatusOK, "DB接続しました")
	}
}

type (
	Response struct {
		Int64  int64  `json:"int64"`
		String string `json:"string"`
		World  *Item  `json:"world"`
	}

	Item struct {
		Text string `json:"text"`
	}
)

// @title go_learn API
// @version 1.0
// @description go_learn API
// @host localhost:1323
// @BasePath /
// @schemes http

func main() {
	e := echo.New()
	e.GET("/", connect)

	e.POST("/users", controller.CreateUser)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
