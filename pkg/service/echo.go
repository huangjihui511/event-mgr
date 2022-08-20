package service

import (
	"bytes"
	"html/template"
	"huangjihui511/event-mgr/pkg/event"
	"huangjihui511/event-mgr/pkg/utils"
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
		event.DashboardData.Lock()
		event.DashboardData.Time = utils.TimeNow()
		t.Execute(buf, event.DashboardData)
		event.DashboardData.Unlock()
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
