package main

import (
	"flag"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	hotReload := flag.Bool("reload", false, "Enable hot-reload")
	flag.Parse()

	server := echo.New()
	server.Use(middleware.Logger())
	server.Renderer = newRenderer()
	server.Static("/static", "public/static")

	if *hotReload {
		hotReloadStart(server)
	}

	server.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", struct{ Reload bool }{Reload: *hotReload})
	})

	server.Logger.Fatal(server.Start(":8080"))
}
