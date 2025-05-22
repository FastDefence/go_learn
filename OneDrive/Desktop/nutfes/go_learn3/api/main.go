package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	//echoを起動
	e := echo.New()
	//path指定して処理を実行させる
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	//ポートを開く、docker-compose.ymlと同じポート番号にする
	e.Logger.Fatal(e.Start(":1323"))
}
