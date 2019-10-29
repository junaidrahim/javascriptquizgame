package main

import (
	// "net/http"
	// "github.com/labstack/echo"
	"javascriptquizgame/parser"
	"javascriptquizgame/updater"
	"fmt"
)

func main() {
	questions := parser.GetQuestions()
	fmt.Println(questions[10])
	updater.ConnectDB()
	// e := echo.New()

	// e.GET("/", func (c echo.Context) error {
	// 	return c.String(http.StatusOK ,"I am javascriptquizgame")
	// })
	// e.Logger.Fatal(e.Start(":1200"))
}
