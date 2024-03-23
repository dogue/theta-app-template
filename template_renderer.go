package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Templates struct {
	template *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func newRenderer() *Templates {
	return &Templates{
		template: template.Must(template.ParseGlob("public/views/*.html")),
	}
}
