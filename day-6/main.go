package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	type Item struct {
		Success 	bool	`json:"success"`
		Data        interface{} `json:"data"`
	}

	type User struct {
		Name 	string	`json:"name"`
	}
	// e := echo.New()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c, echo.Context) error {
	return c.JSON(http.StatusOK,Item){
		Success:	true,
		Data:		c.Param(name: "name"), 		
	}
}