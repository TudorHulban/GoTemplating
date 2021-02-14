package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

// AppConfiguration Structure holding application configuration.
type AppConfiguration struct {
	TemplatesFolder string
	RenderFolder    string
	l               *log.LogInfo
}

func defaultConfiguration() (*AppConfiguration, error) {
	executableFolder, err := os.Getwd()
	if err != nil {
		return nil, errors.WithMessage(err, "issues when creating default configuration")
	}

	result := &AppConfiguration{
		TemplatesFolder: ".." + executableFolder + "/static/assets",
		RenderFolder:    ".." + executableFolder + "/static",
		l:               log.New(log.DEBUG, os.Stdout, true),
	}

	return result, saveConfiguration(result)
}

// saveConfiguration Helper saves configuration if one does not have a file for configuring the app.
func saveConfiguration(cfg *AppConfiguration) error {
	file, errUnmar := json.MarshalIndent(cfg, "", " ")
	if errUnmar != nil {
		return errors.WithMessage(errUnmar, "could not unmarshal configuration")
	}

	return ioutil.WriteFile("default_configuration", file, 0644)
}

// NewConfiguration Constructor for application configuration.
func NewConfiguration(importFrom string, logLevel int) (*AppConfiguration, error) {
	if importFrom == "" {
		return defaultConfiguration()
	}

	data, errRead := ioutil.ReadFile(importFrom)
	if errRead != nil {
		return nil, errors.WithMessagef(errRead, "issues when loading blog articles in file %s", importFrom)
	}

	var result struct {
		Template string
		Render   string
	}

	errUnmar := json.Unmarshal(data, &result)
	if errUnmar != nil {
		return nil, errors.WithMessage(errUnmar, "issues when unmarshaling configuration data")
	}

	return &AppConfiguration{
		TemplatesFolder: result.Template,
		RenderFolder:    result.Render,
		l:               log.New(logLevel, os.Stdout, true),
	}, nil
}
