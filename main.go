package main

import (
	"html/template"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _db *gorm.DB
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
	var errOpen error
	_db, errOpen = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if errOpen != nil {
		log.Fatalf("could not connect to database")
	}

	gin.SetMode(gin.ReleaseMode)

	app := gin.New()
	app.Use(gin.Recovery())
	app.RedirectTrailingSlash = true
	app.HandleMethodNotAllowed = false

	app.GET("/", handler)
	log.Fatal(app.Run("0.0.0.0:8080"))
}
