package articles

// Article model concentrates state and methods for interacting with an article.
// Article represents blog entry to help conversion.
type Article struct {
	Created                   uint64   // UNIX time seconds
	LastUpdated               uint64   // UNIX time seconds
	CODE                      string   `json:"code"`
	Name                      string   `json:"name"`
	AUthor                    string   `json:"author"`
	Content                   string   `json:"content"`
	HTMLTemplate              string   `json:"html"`
	RelatedProductsSKUs       []uint64 `json:"SKUs"`
	RelatedProductsCategories []string `json:"categories"`
}