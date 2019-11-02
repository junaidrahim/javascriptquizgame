package main

import (
	"fmt"
	"javascriptquizgame/database"
)

func main() {
	status := database.UpdateDB()
	fmt.Println(status)

	// e := echo.New()

	// e.GET("/", func (c echo.Context) error {
	// 	return c.String(http.StatusOK ,"I am javascriptquizgame")
	// })
	// e.Logger.Fatal(e.Start(":1200"))
}
