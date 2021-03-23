package main

import (
	"log"
	"os"

	"github.com/TudorHulban/GoTemplating/cmd/config"
	"github.com/TudorHulban/GoTemplating/pkg/httpserve"
)

func main() {
	cfg, errCfg := config.NewConfiguration("", 3)
	if errCfg != nil {
		log.Println(errCfg)
		os.Exit(1)
	}

	cfg.L.Print(*cfg)

	// render landing page

	// start HTTP server
	c := httpserve.Cfg{
		ListenPort:         8008,
		StaticAssetsFolder: "../renderedassets",
	}
	http, errStart := httpserve.NewHTTPServer(c)
	if errStart != nil {
		log.Println(errStart)
		os.Exit(1)
	}

	http.Start()
}
