package main

import (
	"time"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/labstack/echo/v4"
)

func hotReloadStart(server *echo.Echo) {
	s := sse.NewServer(nil)

	server.GET("/events/reload", func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		s.ServeHTTP(res, req)
		return nil
	})

	go keepAlive(s)
}

func keepAlive(s *sse.Server) {
	for {
		s.SendMessage("/events/reload", sse.SimpleMessage("keep-alive"))
		time.Sleep(60 * time.Second)
	}
}
