package app

import (
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
	l                 *log.LogInfo
}

func NewApp(cfg config.AppConfiguration) (*App, error) {
	return nil, nil
}
