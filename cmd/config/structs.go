package config

import (
	"github.com/TudorHulban/GoTemplating/internal/articles"
	"github.com/TudorHulban/GoTemplating/internal/products"
	"github.com/rs/zerolog"
)

type SiteInfo struct {
	FaviconImagePath string
	SiteLogoPath     string
	RenderFolder     string
}

// HTMLPageTemplates Consolidates HTML page templates.
// All templates shold share same containing folder and fields should be file names only.
type HTMLPageTemplates struct {
	ContainingFolder string
	PageShell        string
	Head             string
	Meta             string
	Header           string
	Body             string
	// Section , Aside string
	Article string
	Footer  string
}

// AppConfiguration Structure holding application configuration.
type AppConfiguration struct {
	SiteInfo
	HTMLPageTemplates

	SaveToConfigFile string
	L                zerolog.Logger
}

type App struct {
	SiteInfo

	Templates         map[TemplateName]TemplateContents
	BlogArticles      []articles.Article
	ECommerceProducts []products.Product
	L                 zerolog.Logger
}
