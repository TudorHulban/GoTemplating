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

func NewBlog(importFrom string) (*Blog, error) {
	data, errRead := ioutil.ReadFile(importFrom)
	if errRead != nil {
		return nil, errors.WithMessagef(errRead, "issues when loading blog articles in file %s", importFrom)
	}

	var state []articles.Article

	errUnmar := json.Unmarshal(data, &state)
	if errUnmar != nil {
		return nil, errors.WithMessage(errUnmar, "issues when unmarshaling blog data")
	}

	if len(state) == 0 {
		return nil, errors.New("no articles to batch import")
	}

	for _, art := range state {
		if errValid := articles.ValidateArticle(art); errValid != nil {
			return nil, errors.WithMessagef(errValid, "could not validate article %s", art)
		}
	}

	return &Blog{
		Articles: state,
	}, nil
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
