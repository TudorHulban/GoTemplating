package blogfile

import (
	"encoding/json"
	"io/ioutil"

	"github.com/TudorHulban/GoTemplating/internal/articles"
	"github.com/pkg/errors"
)

type Blog struct {
	Articles []articles.Article
}

// NewBlog Constructor, takes a list of file names and imports them.
// The file names should point to JSON files containing article data.
func NewBlog(importFiles []string) (*Blog, error) {
	result := Blog{
		Articles: []articles{},
	}

	for _, f := range importFiles {
		a, errLoad := result.loadArticle(f)
		if errLoad != nil {
			return nil, errors.WithMessagef(errLoad, "could not load article from %s", f)
		}

		result.Articles = append(result.Articles, a)
	}

	if len(result.Articles) == 0 {
		return nil, errors.New("no articles to batch import")
	}

	return result, nil
}

// loadArticle Loads article from file, performs article validation.
func (b *Blog) loadArticle(loadFrom string) (articles.Article, error) {
	data, errRead := ioutil.ReadFile(loadFrom)
	if errRead != nil {
		return nil, errors.WithMessagef(errRead, "issues when loading blog article in file %s", loadFrom)
	}

	var result articles.Article

	errUnmar := json.Unmarshal(data, &result)
	if errUnmar != nil {
		return nil, errors.WithMessage(errUnmar, "issues when unmarshaling blog article data")
	}

	if errValid := result.ValidateArticle(); errValid != nil {
		return nil, errValid
	}

	return result, nil
}

func (b *Blog) SaveArticles(saveTo string) error {

}

func (b *Blog) saveArticle(a articles.Article) error {

}

// AddArticle Method to be used when adding articles as it offers input validation.
func (b *Blog) AddArticle(a articles.Article) error {
	if errValid := articles.ValidateArticle(a); errValid != nil {
		return errors.WithMessagef(errValid, "could not validate article %s", a)
	}

	b.Articles = append(b.Articles, a)
	return nil
}

// GetArticles Method satisfies articles interface.
func (b *Blog) GetArticles() ([]articles.Article, error) {
	return b.Articles, nil
}

// GetArticle Method satisfies articles interface.
func (b *Blog) GetArticle(code string) (*articles.Article, error) {
	for _, v := range b.Articles {
		if code == v.CODE {
			return &v, nil
		}
	}

	return nil, errors.WithMessage(nil, "no articles found")
}
