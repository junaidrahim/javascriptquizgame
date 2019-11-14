package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"javascriptquizgame/database"
	"net/http"
)

// TODO : set error handling if the database is not started

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

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
