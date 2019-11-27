package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"javascriptquizgame/database"
	"javascriptquizgame/logger"
	"net/http"
	"fmt"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	logger.WriteLog("\n\n")
	logger.WriteLog("Starting Database")

	e.GET("/api/getQuestions", func(c echo.Context) error {
		q := database.GetQuestions()
		if q == nil {
			// Database error occurred
			return c.String(http.StatusInternalServerError, "Internal Error Occurred")
		}
		return c.JSON(http.StatusOK, q)
	})

	e.GET("/api/updateDatabase", func(c echo.Context) error {
		fmt.Println("Starting Database Update")
		logger.WriteLog("Starting Database Update")

		status := database.UpdateDB()

		if !status {
			fmt.Println("Database Error Occurred")
			logger.WriteLog("Database Error Occurred")
			return c.String(http.StatusInternalServerError, "Error")
		} else {
			fmt.Println("Updated Database Successfully")
			logger.WriteLog("Updated Database Successfully")
			return c.String(http.StatusOK, "Successfully Updated")
		}
	})

	e.Logger.Fatal(e.Start(":1200"))
}
