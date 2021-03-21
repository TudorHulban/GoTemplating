package main

import (
	"log"
	"os"

	"github.com/TudorHulban/GoTemplating/internal/app"
	"github.com/TudorHulban/GoTemplating/pkg/httpserve"
)

const configurationFile = "../appconfiguration/cfg.json"

func main() {
	cfg, errCfg := app.NewConfiguration("", 3)
	if errCfg != nil {
		log.Println(errCfg)
		os.Exit(1)
	}

	cfg.L.Print(*cfg)

	// render landing page

	// start HTTP server
	http := httpserve.NewHTTPServer()
	http.Start()
}
