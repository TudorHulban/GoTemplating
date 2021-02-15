package articles

// Article model concentrates state and methods for interacting with an article.
// Article represents blog entry to help conversion.
type Article struct {
	IsVisible                 bool
	Created                   uint64   // UNIX time seconds
	LastUpdated               uint64   // UNIX time seconds
	CODE                      string   `json:"code"`
	Name                      string   `json:"name"`
	Author                    string   `json:"author"`
	Content                   string   `json:"content"`
	HTMLTemplate              string   `json:"html"`
	FeaturedImagePath         string   `json:"featuredimage"`
	RelatedProductsSKUs       []uint64 `json:"SKUs"`
	RelatedProductsCategories []string `json:"categories"`
}

// ValidateArticle Validates article. To be moved in articles package for all implementations to benefit.
func ValidateArticle(a Article) error {
	return nil
}
