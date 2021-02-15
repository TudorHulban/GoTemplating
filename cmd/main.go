package main

import (
	"log"
	"os"

	"github.com/TudorHulban/GoTemplating/internal/config"
)

const configurationFile = "../appconfiguration/cfg.json"

func main() {
	cfg, errCfg := config.NewConfiguration("", 3)
	if errCfg != nil {
		log.Println(errCfg)
		os.Exit(1)
	}

	cfg.L.Print(*cfg)

	// render landing page
}
