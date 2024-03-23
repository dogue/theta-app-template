package main

import (
	"flag"
	"html/template"
	"io"
	"time"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	hotReload := flag.Bool("reload", false, "Enable hot-reload")
	flag.Parse()

	server := echo.New()
	server.Use(middleware.Logger())
	server.Renderer = &Template{templates: template.Must(template.ParseGlob("public/views/*.html"))}
	server.Static("/static", "public/static")

	server.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", struct{ Reload bool }{Reload: *hotReload})
	})

	if *hotReload {
		// hot-reload
		s := sse.NewServer(nil)
		defer s.Shutdown()

		server.GET("/events/reload", func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			s.ServeHTTP(res, req)
			return nil
		})

		go func() {
			for {
				s.SendMessage("/events/reload", sse.SimpleMessage("keep-alive"))
				time.Sleep(60 * time.Second)
			}
		}()
	}

	server.Logger.Fatal(server.Start(":8080"))
}
