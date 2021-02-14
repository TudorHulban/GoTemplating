package main

import (
	"log"
	"os"
)

const configurationFile = "../appconfiguration/cfg.json"

func main() {
	cfg, errCfg := NewConfiguration("", 3)
	if errCfg != nil {
		log.Println(errCfg)
		os.Exit(1)
	}

	cfg.l.Print(*cfg)
}
