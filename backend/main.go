package main

import (
	"github.com/labstack/echo"
	"javascriptquizgame/database"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/api/getQuestions", func (c echo.Context) error {
		q := database.GetQuestions()
		return c.JSON(http.StatusOK, q)
	})

	e.GET("/api/updateDatabase", func (c echo.Context) error {
		status := database.UpdateDB()
		if !status {
			return c.String(http.StatusInternalServerError, "Error")
		} else {
			return c.String(http.StatusOK, "Successfully Updated")
		}
	})
	
	e.Logger.Fatal(e.Start(":1200"))
}
