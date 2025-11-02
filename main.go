package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type requestBody struct {
	Task string `json:"task"`
}

var task string

func createHandler(c echo.Context) error {

	var body requestBody

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Error": "Bad request"})
	}

	task = body.Task

	return c.JSON(http.StatusAccepted, map[string]string{"saved": task})
}

func readHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello +"+task)
}

func main() {	
	e := echo.New()

	e.POST("/calculation", createHandler)
	e.GET("/result", readHandler)

	e.Start("localhost:8080")

}
