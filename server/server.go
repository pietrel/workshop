package server

import (
	"github.com/labstack/echo/v4"
	"second/config"
	"second/handlers"
)

type Server struct {
	echo     *echo.Echo
	port     string
	handlers *handlers.Handlers
}

func NewServer(handlers *handlers.Handlers) *Server {
	e := echo.New()
	e.HideBanner = true
	server := &Server{
		echo:     e,
		port:     config.GetSrvConfig(),
		handlers: handlers,
	}
	server.Routes()

	return server
}

func (server Server) Routes() {
	server.echo.POST("/users", server.handlers.SaveUser)
	server.echo.GET("/users/:id", server.handlers.GetUser)
	server.echo.GET("/users", server.handlers.GetUsers)
	server.echo.PUT("/users/:id", server.handlers.UpdateUser)
	server.echo.DELETE("/users/:id", server.handlers.DeleteUser)
}

func (server Server) Run() {
	server.echo.Logger.Fatal(server.echo.Start(server.port))
}
