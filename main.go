package main

import (
	"fmt"
	_ "l-gomock-swaggo/docs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoswagger "github.com/swaggo/echo-swagger"
)

// @title        sample API
// @version      1.0
// @description  This is a sample
// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
// @host         localhost:8080
// @BasePath     /
func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoswagger.WrapHandler)
	e.GET("/hello", hello)

	port := "8080"
	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}

// Hello
//
// @Summary     Hello
// @Description Hello Go
// @Accept      json
// @Produce     json
// @Success     200 {object} string
// @Router      /hello [get]
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Go!")
}
