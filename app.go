package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Criaturitas del Universo!!! <3")
	})

	e.GET("/betis", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<p style='text-align:center;'><a href='https://en.realbetisbalompie.es' target='_blank'><img src='https://upload.wikimedia.org/wikipedia/en/thumb/1/13/Real_betis_logo.svg/1024px-Real_betis_logo.svg.png'></a></p>")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
