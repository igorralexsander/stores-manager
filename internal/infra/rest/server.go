package rest

import (
	"fmt"
	"github.com/igorralexsander/stores-manager/internal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Server struct {
	app *internal.App
}

func NewServer(app *internal.App) *Server {
	return &Server{
		app: app,
	}
}

func (s *Server) CreateHttpServer() *echo.Echo {
	e := echo.New()
	s.app.RegisterRoutes(e)
	return e
}

func (s *Server) Start(server *echo.Echo, host string) {
	log.Info(fmt.Sprintf("Starting http server at http://%s", host))
	log.Warn(server.Start(host), "Stopping application....")

}
