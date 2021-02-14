package articles

import "github.com/pkg/errors"

type Blog struct {
	Articles []Article
}

func NewBlog() (*Blog, error) {
	return &Blog{
		Articles: []Article{},
	}, nil
}

func (b *Blog) AddArticle(a Article) error {
	// TODO: add article validation

	b.Articles = append(b.Articles, a)
	return nil
}

func (b *Blog) GetArticles() ([]Article, error) {
	return b.Articles, nil
}

func (b *Blog) GetArticle(code string) (*Article, error) {
	for _, v := range b.Articles {
		if code == v.CODE {
			return &v, nil
		}
	}

	return nil, errors.WithMessage(nil, "no articles found")
}
