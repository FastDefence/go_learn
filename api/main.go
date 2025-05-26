package main

import (
	"net/http"

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

func main() {
	//echoを起動
	e := echo.New()
	//path指定して処理を実行させる
	e.GET("/", connect)
	//ポートを開く、docker-compose.ymlと同じポート番号にする
	e.Logger.Fatal(e.Start(":1323"))
}
