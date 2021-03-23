package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// NewConfiguration Constructor for application configuration.
// TODO: consider moving to options once the validation criterias emerge.
func NewConfiguration(importFrom string, logLevel zerolog.Level) (*AppConfiguration, error) {
	if importFrom == "" {
		return defaultConfiguration()
	}

	data, errRead := ioutil.ReadFile(importFrom)
	if errRead != nil {
		return nil, errors.WithMessagef(errRead, "issues when loading blog articles in file %s", importFrom)
	}

	var result struct {
		SiteInfo
		HTMLPageTemplates
	}

	errUnmar := json.Unmarshal(data, &result)
	if errUnmar != nil {
		return nil, errors.WithMessage(errUnmar, "issues when unmarshaling configuration data")
	}

	return &AppConfiguration{
		SiteInfo:          result.SiteInfo,
		HTMLPageTemplates: result.HTMLPageTemplates,
		L:                 zerolog.New(os.Stderr).With().Timestamp().Logger().Level(logLevel),
	}, nil
}

// saveConfiguration Helper saves configuration if one does not have a file for configuring the app.
func saveConfiguration(cfg *AppConfiguration) error {
	file, errUnmar := json.MarshalIndent(cfg, "", " ")
	if errUnmar != nil {
		return errors.WithMessage(errUnmar, "could not unmarshal configuration")
	}

	if cfg.SaveToConfigFile == "" {
		return ioutil.WriteFile(defaultAppConfigurationFileName, file, 0644)
	}

	return ioutil.WriteFile(cfg.SaveToConfigFile, file, 0644)
}
