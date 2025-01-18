package app

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/wehw93/test_task_Junior_go_developer/internal/app/endpoint"
	"github.com/wehw93/test_task_Junior_go_developer/internal/app/mw"
	"github.com/wehw93/test_task_Junior_go_developer/internal/app/service"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()
	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck)
	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server runnings")
	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
