package app

import (
	"github.com/TudorHulban/GoTemplating/internal/articles"
	"github.com/TudorHulban/GoTemplating/internal/products"
	"github.com/TudorHulban/log"
)

type TemplateName string
type TemplateContents string

type App struct {
	SiteInfo

	Templates         map[TemplateName]TemplateContents
	BlogArticles      []articles.Article
	ECommerceProducts []products.Product
	L                 *log.LogInfo
}

func NewApp(cfg AppConfiguration) (*App, error) {
	return nil, nil
}
