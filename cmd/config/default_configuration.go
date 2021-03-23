package config

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Option func(cfg *AppConfiguration) error

func defaultConfiguration(options ...Option) (*AppConfiguration, error) {
	executableFolder, err := os.Getwd()
	if err != nil {
		return nil, errors.WithMessage(err, "issues when creating default configuration")
	}

	renderToFolder := executableFolder + defaultRenderArticlesFolder

	result := &AppConfiguration{
		SiteInfo: SiteInfo{
			RenderArticlesFolder: renderToFolder,
		},

		AppConfigFile: defaultAppConfigurationFileName,

		HTMLPageTemplates: HTMLPageTemplates{
			ContainingFolder: ".." + executableFolder + "/static/assets",
			PageShell:        "01_page_shell.gohtml",
			Head:             "02_head.gohtml",
			Meta:             "03_meta.gohtml",
			Header:           "04_header.gohtml",
			Body:             "05_body.gohtml",
			Article:          "06_article.gohtml",
			Footer:           "07_footer.gohtml",
		},

		L: zerolog.New(os.Stderr).With().Caller().Timestamp().Logger().Level(zerolog.DebugLevel),
	}

	// moved below initialization in order to use the logger
	result.L.Print(renderToFolder)

	if _, err := os.Stat(renderToFolder); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(renderToFolder, os.ModePerm)
		}
	}

	return result, saveConfiguration(result)
}
