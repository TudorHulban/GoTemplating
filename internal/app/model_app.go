package app

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/pkg/errors"

	"github.com/TudorHulban/GoTemplating/internal/articles"
	"github.com/TudorHulban/GoTemplating/internal/config"
	"github.com/TudorHulban/GoTemplating/internal/products"
	"github.com/TudorHulban/log"
)

type TemplateName string
type TemplateContents string

type App struct {
	config.SiteInfo
	Templates         map[TemplateName]TemplateContents
	BlogArticles      []articles.Article
	ECommerceProducts []products.Product
	L                 *log.LogInfo
}

func NewApp(cfg config.AppConfiguration) (*App, error) {
	return nil, nil
}

func defaultArticles() []articles.Article {
	return []articles.Article{articles.Article{
		IsVisible:   true,
		Created:     uint64(time.Now().Unix()),
		LastUpdated: 0,
		CODE:        "ART01",
		Name:        "Default Article",
		Author:      "Default Author",
		Content:     "xxxxxxxxxxxxxxxxxxx",
	},
	}
}

func saveArticles(art []articles.Article) error {
	if len(art) == 0 {
		errors.New("no articles to save")
	}

	file, errUnmar := json.MarshalIndent(art, "", " ")
	if errUnmar != nil {
		return errors.WithMessage(errUnmar, "could not unmarshal configuration")
	}

	return ioutil.WriteFile("default_configuration", file, 0644)
}

func loadArticles(importFrom string) ([]articles.Article, error) {
	if importFrom == "" {
		return nil, errors.New("did not provide articles file to import from")
	}

}
