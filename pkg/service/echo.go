package service

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func startEcho() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		t := template.Must(template.ParseGlob("./template/index.tmp"))
		b := make([]byte, 0)
		buf := bytes.NewBuffer(b)
		t.Execute(buf, nil)
		return c.HTML(http.StatusOK, string(buf.Bytes()))
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
