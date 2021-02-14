package articles

type IArticles interface {
	GetArticles() ([]Article, error)
	GetArticle(code string) (Article, error)
	GetRelatedArticles(sku uint64, howMany uint8) (Article, error)
	GetCategoryArticles(category string, howMany uint8) (Article, error)
}
