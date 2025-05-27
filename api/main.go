package main

import (
	"net/http"

	_ "github.com/FastDefence/go_learn/api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"

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
	//echoを起動
	e := echo.New()
	//path指定して処理を実行させる
	e.GET("/", connect)
	//ポートを開く、docker-compose.ymlと同じポート番号にする
	e.Logger.Fatal(e.Start(":1323"))
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
