package articles

import (
	"time"
)

func DefaultArticles() []Article {
	a1, a2 := "ART01", "ART02"

	return []Article{Article{
		IsVisible:   true,
		Created:     uint64(time.Now().Unix()),
		LastUpdated: 0,
		CODE:        a1,
		SaveToFile:  a1 + ".art",
		Name:        "Default Article",
		Author:      "Default Author",
		Content:     "xxxxxxxxxxxxxxxxxxx",
	},
		Article{
			IsVisible:   true,
			Created:     uint64(time.Now().Unix()),
			LastUpdated: 0,
			CODE:        a2,
			SaveToFile:  a2 + ".art",
			Name:        "Default Article",
			Author:      "Default Author",
			Content:     "xxxxxxxxxxxxxxxxxxx",
		},
	}
}
