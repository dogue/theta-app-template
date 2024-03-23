package main

import (
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
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = &Template{templates: template.Must(template.ParseGlob("public/views/*.html"))}
	e.Static("/static", "public/static")
	s := sse.NewServer(nil)
	defer s.Shutdown()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})

	// hot-reload
	e.Any("/events/:channel", func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		s.ServeHTTP(res, req)
		return nil
	})

	go func() {
		for {
			s.SendMessage("/events/reload", sse.SimpleMessage("keep-alive"))
			time.Sleep(30 * time.Second)
		}
	}()

	e.Logger.Fatal(e.Start(":8080"))
}
