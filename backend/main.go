package main

import (
	// "net/http"
	// "github.com/labstack/echo"
	//"fmt"
	"javascriptquizgame/parser"
)

func main() {
	parser.ParseGlob(parser.GetData())

	// e := echo.New()

	// e.GET("/", func (c echo.Context) error {
	// 	return c.String(http.StatusOK ,"I am javascriptquizgame")
	// })
	// e.Logger.Fatal(e.Start(":1200"))
}
