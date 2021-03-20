package articles

// Article model concentrates state and methods for interacting with an article.
// Article represents blog entry to help conversion.
type Article struct {
	IsVisible                 bool     `json:"visible"`
	Created                   uint64   `json:"created"` // UNIX time seconds
	LastUpdated               uint64   `json:"updated"` // UNIX time seconds
	SaveToFile                string   `json:"file"`
	CODE                      string   `json:"code"`
	Name                      string   `json:"name"`
	Author                    string   `json:"author"`
	Content                   string   `json:"content"`
	HTMLTemplateFile          string   `json:"html"`
	FeaturedImagePath         string   `json:"featuredimage"`
	RelatedProductsSKUs       []uint64 `json:"relatedsku"`
	RelatedProductsCategories []string `json:"relatedcateg"`
}

// ValidateArticle Validates article. To be moved in articles package for all implementations to benefit.
func (a Article) ValidateArticle() error {
	return nil
}
