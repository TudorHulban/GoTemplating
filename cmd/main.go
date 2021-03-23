package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/TudorHulban/GoTemplating/cmd/config"
	"github.com/TudorHulban/GoTemplating/internal/blog/blogfile"
	"github.com/TudorHulban/GoTemplating/pkg/httpserve"
)

func main() {
	cfg, errCfg := config.NewConfiguration("", 3)
	if errCfg != nil {
		log.Println(errCfg)
		os.Exit(1)
	}

	articles, errGet := getFiles(cfg.RenderArticlesFolder, config.ExtensionArticleFile)
	if errGet != nil {
		cfg.L.Print(errGet)
		os.Exit(2)
	}

	if len(articles) == 0 {
		cfg.L.Print("could not find article files")
		os.Exit(3)
	}

	blog, errBlog := blogfile.NewBlogFromFiles(cfg.L, articles...)
	if errBlog != nil {
		cfg.L.Print(errBlog)
		os.Exit(4)
	}

	// render article pages
	errRender := blog.RenderArticles()
	if errRender != nil {
		cfg.L.Print(errRender)
		os.Exit(5)
	}

	// start HTTP server
	c := httpserve.Cfg{
		ListenPort:         8008,
		StaticAssetsFolder: "../renderedassets",
	}

	http, errStart := httpserve.NewHTTPServer(c)
	if errStart != nil {
		log.Println(errStart)
		os.Exit(6)
	}

	http.Start()
}

func getFiles(fromFolder, withExtension string) ([]string, error) {
	// TODO: improve validation

	files, err := ioutil.ReadDir(fromFolder)
	if err != nil {
		return []string{}, err
	}

	result := []string{}

	for _, v := range files {
		if v.IsDir() == false {
			result = append(result, v.Name())
		}
	}

	return result, nil
}
