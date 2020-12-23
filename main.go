package main

import (
	"html/template"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Product struct {
	ID                 uint
	SKU                string
	Category           string
	Slug               string
	Name               string
	Description        string
	SEOMetaTitle       string
	SEOMetaDescription string
	Quantity           uint
	PriceCents         float32
}

type ProductImage struct {
	ProductID uint
	MetaALT   string // holds alt for image
	Path      string
}

var templ *template.Template

// Used code as per:
// https://stackoverflow.com/questions/45828142/golang-multi-templates-caching
func prepTemplate() (*template.Template, error) {
	parsedTemplate, errParse := template.ParseFiles("index.gohtml")
	if errParse != nil {
		return nil, errors.Wrap(errParse, "could not parse template")
	}

	var errMust error
	return template.Must(parsedTemplate, errMust), errMust
}

func render(w io.Writer, p Product) error {
	if templ == nil {
		var errPrep error
		templ, errPrep = prepTemplate()
		if errPrep != nil {
			log.Fatal("could not prepare template")
		}
	}

	errExec := templ.Execute(w, p)
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
