package main

import (
	user "echo-api/controller"
	"echo-api/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	ctrl := user.Controller{}
	e.GET("/users", ctrl.Index)
	e.GET("/users/:id", ctrl.Show)
	e.Logger.Fatal(e.Start(":1323"))
	db.Close()
}
