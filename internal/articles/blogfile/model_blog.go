package blogfile

import (
	"encoding/json"
	"io/ioutil"

	"github.com/TudorHulban/GoTemplating/internal/articles"
	"github.com/pkg/errors"
)

var _ articles.IBlog = (*Blog)(nil)

type Blog struct {
	Data []articles.Article
}

func NewBlogFromArticles(art ...articles.Article) (*Blog, error) {
	result := Blog{
		Data: []articles.Article{},
	}

	if len(art) == 0 {
		return nil, errors.New("no articles provided")
	}

	for _, a := range art {
		if err := result.AddArticle(a); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

// NewBlog Constructor, takes a list of file names and imports them.
// The file names should point to JSON files containing article data.
func NewBlogFromFiles(importFiles ...string) (*Blog, error) {
	result := Blog{
		Data: []articles.Article{},
	}

	if len(importFiles) == 0 {
		return nil, errors.New("no import files provided")
	}

	for _, f := range importFiles {
		a, errLoad := result.loadArticle(f)
		if errLoad != nil {
			return nil, errors.WithMessagef(errLoad, "could not load article from %s", f)
		}

		result.Data = append(result.Data, *a)
	}

	if len(result.Data) == 0 {
		return nil, errors.New("no articles to batch import")
	}

	return &result, nil
}

// loadArticle Loads article from file, performs article validation.
func (b *Blog) loadArticle(loadFrom string) (*articles.Article, error) {
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

	return &result, nil
}

// AddArticle Method to be used when adding articles as it offers input validation.
func (b *Blog) AddArticle(a articles.Article) error {
	if errValid := articles.Article.ValidateArticle(a); errValid != nil {
		return errors.WithMessagef(errValid, "could not validate article %v", a)
	}

	b.Data = append(b.Data, a)
	return nil
}

// GetArticles Method satisfies articles interface.
func (b *Blog) GetArticles() ([]articles.Article, error) {
	return b.Data, nil
}

// GetArticle Method satisfies articles interface.
func (b *Blog) GetArticle(code string) (*articles.Article, error) {
	for _, v := range b.Data {
		if code == v.CODE {
			return &v, nil
		}
	}

	return nil, errors.WithMessage(nil, "no articles found")
}

func (b *Blog) SaveBlogArticles() error {
	for _, a := range b.Data {
		if err := b.saveArticle(a); err != nil {
			return err
		}
	}

	return nil
}

func (b *Blog) saveArticle(a articles.Article) error {
	byteArticle, errMar := json.MarshalIndent(a, "", " ")
	if errMar != nil {
		return errMar
	}

	return ioutil.WriteFile(a.SaveToFile, byteArticle, 0644)
}
