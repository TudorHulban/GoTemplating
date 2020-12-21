package main

import (
	"html/template"
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
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

func handler(c *fiber.Ctx) error {
	p := Product{
		Name:       "Apple",
		Quantity:   400,
		PriceCents: 150,
	}

	return render(c.Response().BodyWriter(), p)
}

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})

	app.Get("/", handler)
	log.Fatal(app.Listen(":3000"))
}
