package main

import (
	"html/template"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Product struct {
	Name       string
	Quantity   uint
	PriceCents float32
}

func render(w io.Writer, p Product) error {
	parsedTemplate, errParse := template.ParseFiles("index.html")
	if errParse != nil {
		return errors.Wrap(errParse, "could not parse template")
	}

	errExec := parsedTemplate.Execute(w, p)
	if errExec != nil {
		return errors.Wrap(errExec, "could not execute template")
	}

	return nil
}

func handler(c *gin.Context) {
	p := Product{
		Name:       "Apple",
		Quantity:   400,
		PriceCents: 150,
	}

	render(c.Writer, p)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	app := gin.New()
	app.Use(gin.Recovery())
	app.RedirectTrailingSlash = true
	app.HandleMethodNotAllowed = false

	app.GET("/", handler)
	log.Fatal(app.Run("0.0.0.0:8080"))
}
